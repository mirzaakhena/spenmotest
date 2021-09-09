package model

import (
  "fmt"
  "time"
)

// Spend is a service that will reduce the amount of wallet balance for specific card in certain time
func (w *Wallet) Spend(repo FindLastUserSpendRepo, card *Card, amount Money, now time.Time) (*UserSpendMoney, error) {

  lastUserSpend, err := repo.FindLastUserSpend(w.User, card)
  if err != nil {
    return nil, err
  }

  // never spend on that card yet. This is the first time user spend it
  if lastUserSpend == nil {
    return w.newLimitToSpend(card.LimitAmount, amount, card, now)
  }

  stillPossibleToSpend := false

  if card.LimitDuration == LimitTimeEnumDaily {
    stillPossibleToSpend, err = w.inTheSameDay(lastUserSpend.Date, now)
    if err != nil {
      return nil, err
    }
  } else //

  if card.LimitDuration == LimitTimeEnumWeekly {
    stillPossibleToSpend, err = w.inTheSameWeek(lastUserSpend.Date, now)
    if err != nil {
      return nil, err
    }
  } else //

  if card.LimitDuration == LimitTimeEnumMonthly {
    stillPossibleToSpend, err = w.inTheSameMonth(lastUserSpend.Date, now)
    if err != nil {
      return nil, err
    }
  }

  if stillPossibleToSpend {

    if lastUserSpend.BalanceRemaining == 0 {
      return nil, fmt.Errorf("your balance limit is 0")
    }

    if amount > lastUserSpend.BalanceRemaining {
      return nil, fmt.Errorf("your balance limit is only %1.0f in card %s. not enough to spend %1.0f", lastUserSpend.BalanceRemaining, card.ID, amount)
    }

    newUserSpend := UserSpendMoney{
      User:             w.User,
      Card:             card,
      Amount:           amount,
      BalanceRemaining: lastUserSpend.BalanceRemaining - amount,
      Date:             now,
    }

    w.Balance = w.Balance - amount

    return &newUserSpend, nil

  }

  // it is in the next day
  return w.newLimitToSpend(lastUserSpend.BalanceRemaining, amount, card, now)

}

func (w *Wallet) newLimitToSpend(balanceRemaing Money, amount Money, card *Card, now time.Time) (*UserSpendMoney, error) {

  if amount > w.Balance {
    // TODO for security reason we must be aware to display the current balance
    return nil, fmt.Errorf("your balance is only %1.0f", w.Balance)
  }

  if amount > card.LimitAmount {
    return nil, fmt.Errorf("your card %s remaining balance limit is only %1.0f", card.ID, balanceRemaing)
  }

  newUserSpend := UserSpendMoney{
    User:             w.User,
    Card:             card,
    Amount:           amount,
    BalanceRemaining: card.LimitAmount - amount,
    Date:             now,
  }

  w.Balance = w.Balance - amount

  return &newUserSpend, nil

}

func (w *Wallet) inTheSameDay(lastDate time.Time, now time.Time) (bool, error) {

  err := w.validateNowAsFuture(lastDate, now)
  if err != nil {
    return false, err
  }

  y1, m1, d1 := lastDate.Date()
  y2, m2, d2 := now.Date()

  return y1 == y2 && m1 == m2 && d1 == d2, nil
}

func (w *Wallet) inTheSameWeek(lastDate time.Time, now time.Time) (bool, error) {

  err := w.validateNowAsFuture(lastDate, now)
  if err != nil {
    return false, err
  }

  y1, w1 := lastDate.ISOWeek()
  y2, w2 := lastDate.ISOWeek()

  return y1 == y2 && w1 == w2, nil
}

func (w *Wallet) inTheSameMonth(lastDate time.Time, now time.Time) (bool, error) {

  err := w.validateNowAsFuture(lastDate, now)
  if err != nil {
    return false, err
  }

  y1, m1, _ := lastDate.Date()
  y2, m2, _ := now.Date()

  return y1 == y2 && m1 == m2, nil
}

func (w *Wallet) validateNowAsFuture(lastDate time.Time, now time.Time) error {
  if now.Before(lastDate) {
    return fmt.Errorf("now must be future than lastDate")
  }
  return nil
}
