package exercise3

import "datastructure/exercise3/model"

// Situation is the captured event in some universe event
type Situation struct {
  Wallets   []*model.Wallet
  UserTeams []*model.UserTeam
}

// GetDefaultSituation is the sample of Situation
func GetDefaultSituation() Situation {

  // user has multiple wallets
  // each wallet have multiple card
  // each card has limit amount and duration of use (daily, monthly)
  // user can only spend the money through only one card.

  // user mirza has 2 wallet,
  // first wallet is w1 with total balance 1.000
  // w1 has 2 card, cardone and cardtwo.
  // both card share balance from wallet and has different amount of limit and duration of limit
  //
  // the second wallet is w2 with total balance 20.000
  // this wallet is also belong to a team (teamone) that lead by mirza
  user1 := model.User{ID: "mirza"}

  // This is the mirza's wallet current condition
  wallets := []*model.Wallet{
    {
      User: &user1,
      ID:   "w1", Balance: model.Money(1000),
      Cards: []*model.Card{
        {ID: "cardone", LimitAmount: 300, LimitDuration: model.LimitTimeEnumDaily},
        {ID: "cardtwo", LimitAmount: 100, LimitDuration: model.LimitTimeEnumMonthly},
      },
    },
    {
      User: &user1,
      ID:   "w2", Balance: model.Money(20000),
      Cards: []*model.Card{
        {ID: "cardthree", LimitAmount: 5000, LimitDuration: model.LimitTimeEnumWeekly},
      },
    },
  }

  // let say there is another two users, John and Paul
  // we can assume that they also have a personal wallets with a balance and a cards, but I don't explain it here.
  user2 := model.User{ID: "john"}
  user3 := model.User{ID: "paul"}

  // Mirza, John and Paul is part of teamone
  // teamone has a wallet that manage by mirza
  team1 := model.Team{
    ID:     "teamone",
    Wallet: wallets[1],
  }

  // we link Mirza, John and Paul with this structure.
  // since Mirza who is manage the wallet, then Mirza is an Admin
  // it will be possible for the user team to have more than one admin.
  // role can be also extend to other capability in the future.
  userTeams := []*model.UserTeam{
    {
      Team: &team1,
      User: &user1,
      Role: model.UserTeamRoleAdmin,
    },
    {
      Team: &team1,
      User: &user2,
      Role: model.UserTeamRoleMember,
    },
    {
      Team: &team1,
      User: &user3,
      Role: model.UserTeamRoleMember,
    },
  }

  return Situation{
    Wallets:   wallets,
    UserTeams: userTeams,
  }

}
