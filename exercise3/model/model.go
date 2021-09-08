package model

import (
  "time"
)

type User struct {
  ID Identifier
}

type Wallet struct {
  ID      Identifier
  User    *User
  Balance Money
  Cards   []*Card
}

type Card struct {
  ID            Identifier
  LimitAmount   Money
  LimitDuration LimitTimeEnum
}

type UserSpendMoney struct {
  TransactionID    Identifier // can be used for more detail information in other tables but not explained here
  User             *User
  Card             *Card
  Amount           Money // currently it is not very useful here
  BalanceRemaining Money
  Date             time.Time
}

type Team struct {
  ID     Identifier
  Wallet *Wallet
}

type UserTeam struct {
  User *User
  Team *Team
  Role UserTeamRole
}

type Money float64

type LimitTimeEnum time.Duration

type Identifier string

const (
  LimitTimeEnumDaily   = LimitTimeEnum(time.Hour * 24)
  LimitTimeEnumWeekly  = LimitTimeEnum(time.Hour * 24 * 7)
  LimitTimeEnumMonthly = LimitTimeEnum(time.Hour * 24 * 7 * 30)
)

type UserTeamRole string

const (
  UserTeamRoleAdmin  = UserTeamRole("admin")
  UserTeamRoleMember = UserTeamRole("member")
)

type FindLastUserSpendRepo interface {
  FindLastUserSpend(user *User, card *Card) (*UserSpendMoney, error)
}
