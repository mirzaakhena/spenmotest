package model

import (
  "time"
)

// User is simplified version of User information
// any usefull detail can be add later
type User struct {
  ID Identifier
}

// Wallet hold one user reference annd hold multiple cards
// Balance will be shared to all cards
type Wallet struct {
  ID      Identifier
  User    *User
  Balance Money
  Cards   []*Card
}

// Card information like LimitAmount or LimitDuration is user defined
type Card struct {
  ID            Identifier
  LimitAmount   Money
  LimitDuration LimitTimeEnum
}

// CardSpendHistory is a transaction record to hold the history of user regarding spending money activity on each card.
// BalanceRemaining always less than the Card LimitAmount and will be reset to Card LimitAmount after the LimitDuration
// Date capture the transaction date
// TransactionID can be used to link to other tables such as accounting or order for example
// Amount show how much user spend money in that curret transaction
type CardSpendHistory struct {
  TransactionID    Identifier
  User             *User
  Card             *Card
  Amount           Money // currently it is not very useful here
  BalanceRemaining Money
  Date             time.Time
}

// Team only hold one Wallet reference
// Wallet owned and managed by specific User
type Team struct {
  ID     Identifier
  Wallet *Wallet
}

// UserTeam bind User and team with Role
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

// FindLastCardSpendRepo used by Spend service to find the last CardSpendHistory record
type FindLastCardSpendRepo interface {
  FindLastCardSpend(user *User, card *Card) (*CardSpendHistory, error)
}
