package exercise3

import (
  "datastructure/exercise3/model"
  "time"
)

type normalLastUserSpend struct {
  amount           model.Money
  balanceRemaining model.Money
  lastTime         time.Time
}

func NewLastUserSpendCase(amount, balanceRemaining model.Money, lastTime time.Time) *normalLastUserSpend {
  return &normalLastUserSpend{amount: amount, balanceRemaining: balanceRemaining, lastTime: lastTime}
}

func (l *normalLastUserSpend) FindLastCardSpend(user *model.User, card *model.Card) (*model.CardSpendHistory, error) {
  return &model.CardSpendHistory{
    User:             user,
    Card:             card,
    Amount:           l.amount,
    BalanceRemaining: l.balanceRemaining,
    Date:             l.lastTime,
  }, nil
}

type nilLastUserSpend struct {}

func (l *nilLastUserSpend) FindLastCardSpend(user *model.User, card *model.Card) (*model.CardSpendHistory, error) {
  return nil, nil
}

func NilLastUserSpendCase() *nilLastUserSpend {
  return &nilLastUserSpend{}
}