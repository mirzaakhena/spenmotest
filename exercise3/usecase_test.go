package exercise3

import (
  "datastructure/exercise3/model"
  "github.com/stretchr/testify/assert"
  "testing"
  "time"
)

func TestCase1(t *testing.T) {

  situation := GetDefaultSituation()

  lastTime := time.Now()

  // 5 hour later after now
  now := lastTime.Add(time.Hour * 5)

  // never used the balance
  lastUserSpend := NilLastUserSpendCase()

  // spending 100
  spend, _ := situation.Wallets[0].Spend(lastUserSpend, situation.Wallets[0].Cards[0], 100, now)

  // everything run normaly
  assert.Equal(t, model.Money(100), spend.Amount)
  assert.Equal(t, model.Money(200), spend.BalanceRemaining)
  assert.Equal(t, model.Money(900), situation.Wallets[0].Balance)

}

func TestCase2(t *testing.T) {

  situation := GetDefaultSituation()

  lastTime, _ := time.Parse("2006-01-02 15:04:05", "2021-01-01 13-13-13")

  // 5 hour later after now
  now := lastTime.Add(time.Hour * 5)

  // have been used the balance
  lastUserSpend := NewLastUserSpendCase(20, 200, now)

  // spending more 100
  spend, _ := situation.Wallets[0].Spend(lastUserSpend, situation.Wallets[0].Cards[0], 100, now)

  // everything run normaly
  assert.Equal(t, model.Money(100), spend.Amount)
  assert.Equal(t, model.Money(100), spend.BalanceRemaining)
  assert.Equal(t, model.Money(900), situation.Wallets[0].Balance)

}

func TestCase3(t *testing.T) {

  situation := GetDefaultSituation()

  lastTime, _ := time.Parse("2006-01-02 15:04:05", "2021-01-01 13-13-13")

  // two days after last spending
  now := lastTime.Add(time.Hour * 24 * 2)

  // remaining balance only 50
  lastUserSpend := NewLastUserSpendCase(100, 50, lastTime)

  // spending 100
  spend, _ := situation.Wallets[0].Spend(lastUserSpend, situation.Wallets[0].Cards[0], 100, now)

  // even if remaining balance only 50 but all balance is reseted to card limit amount
  assert.Equal(t, model.Money(100), spend.Amount)
  assert.Equal(t, model.Money(200), spend.BalanceRemaining)
  assert.Equal(t, model.Money(900), situation.Wallets[0].Balance)

}

func TestCase4(t *testing.T) {

  situation := GetDefaultSituation()

  lastTime, _ := time.Parse("2006-01-02 15:04:05", "2021-01-01 13-13-13")

  // 5 hour later after now
  now := lastTime.Add(time.Hour * 5)

  // use all the balance in the card
  lastUserSpend := NewLastUserSpendCase(10, 0, lastTime)

  // spending 100, which is not possible
  _, err := situation.Wallets[0].Spend(lastUserSpend, situation.Wallets[0].Cards[0], 100, now)

  assert.EqualError(t,err, "your balance limit is 0")

}

func TestCase5(t *testing.T) {

  situation := GetDefaultSituation()

  lastTime, _ := time.Parse("2006-01-02 15:04:05", "2021-01-01 13-13-13")

  // 5 hour later after now
  now := lastTime.Add(time.Hour * 5)

  // use all the balance in the card
  lastUserSpend := NewLastUserSpendCase(10, 50, lastTime)

  // spending 60, which is not possible
  _, err := situation.Wallets[0].Spend(lastUserSpend, situation.Wallets[0].Cards[0], 60, now)

  assert.EqualError(t,err, "your balance limit is only 50 in card cardone. not enough to spend 60")

}

func TestCase6(t *testing.T) {

  situation := GetDefaultSituation()

  lastTime := time.Now()

  // 5 hour later after now
  now := lastTime.Add(time.Hour * 5)

  // never used the balance
  lastUserSpend := NilLastUserSpendCase()

  // spending 1000 which is not possible
  _, err := situation.Wallets[0].Spend(lastUserSpend, situation.Wallets[0].Cards[0], 1000, now)

  assert.EqualError(t,err, "your card cardone remaining balance limit is only 300")

}
