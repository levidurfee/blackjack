package strategy

import "git.x6c.co/go/blackjack/pkg/blackjack"

var Pair = StrategyTable{
	// Pair of aces
	2: {
		blackjack.TwoName:   blackjack.Split,
		blackjack.ThreeName: blackjack.Split,
		blackjack.FourName:  blackjack.Split,
		blackjack.FiveName:  blackjack.Split,
		blackjack.SixName:   blackjack.Split,
		blackjack.SevenName: blackjack.Split,
		blackjack.EightName: blackjack.Split,
		blackjack.NineName:  blackjack.Split,
		blackjack.TenName:   blackjack.Split,
		blackjack.JackName:  blackjack.Split,
		blackjack.QueenName: blackjack.Split,
		blackjack.KingName:  blackjack.Split,
		blackjack.AceName:   blackjack.Split,
	},
	// Pair of twos
	4: {
		blackjack.TwoName:   blackjack.Hit,
		blackjack.ThreeName: blackjack.Hit,
		blackjack.FourName:  blackjack.Split,
		blackjack.FiveName:  blackjack.Split,
		blackjack.SixName:   blackjack.Split,
		blackjack.SevenName: blackjack.Split,
		blackjack.EightName: blackjack.Hit,
		blackjack.NineName:  blackjack.Hit,
		blackjack.TenName:   blackjack.Hit,
		blackjack.JackName:  blackjack.Hit,
		blackjack.QueenName: blackjack.Hit,
		blackjack.KingName:  blackjack.Hit,
		blackjack.AceName:   blackjack.Hit,
	},
	6: {
		blackjack.TwoName:   blackjack.Hit,
		blackjack.ThreeName: blackjack.Hit,
		blackjack.FourName:  blackjack.Split,
		blackjack.FiveName:  blackjack.Split,
		blackjack.SixName:   blackjack.Split,
		blackjack.SevenName: blackjack.Split,
		blackjack.EightName: blackjack.Hit,
		blackjack.NineName:  blackjack.Hit,
		blackjack.TenName:   blackjack.Hit,
		blackjack.JackName:  blackjack.Hit,
		blackjack.QueenName: blackjack.Hit,
		blackjack.KingName:  blackjack.Hit,
		blackjack.AceName:   blackjack.Hit,
	},
	// Pair of fours
	8: {
		blackjack.TwoName:   blackjack.Hit,
		blackjack.ThreeName: blackjack.Hit,
		blackjack.FourName:  blackjack.Hit,
		blackjack.FiveName:  blackjack.Hit,
		blackjack.SixName:   blackjack.Hit,
		blackjack.SevenName: blackjack.Hit,
		blackjack.EightName: blackjack.Hit,
		blackjack.NineName:  blackjack.Hit,
		blackjack.TenName:   blackjack.Hit,
		blackjack.JackName:  blackjack.Hit,
		blackjack.QueenName: blackjack.Hit,
		blackjack.KingName:  blackjack.Hit,
		blackjack.AceName:   blackjack.Hit,
	},
	// Pair of fives
	10: {
		blackjack.TwoName:   blackjack.Double,
		blackjack.ThreeName: blackjack.Double,
		blackjack.FourName:  blackjack.Double,
		blackjack.FiveName:  blackjack.Double,
		blackjack.SixName:   blackjack.Double,
		blackjack.SevenName: blackjack.Double,
		blackjack.EightName: blackjack.Double,
		blackjack.NineName:  blackjack.Double,
		blackjack.TenName:   blackjack.Double,
		blackjack.JackName:  blackjack.Double,
		blackjack.QueenName: blackjack.Double,
		blackjack.KingName:  blackjack.Hit,
		blackjack.AceName:   blackjack.Hit,
	},
	// Pair of sixes
	12: {
		blackjack.TwoName:   blackjack.Split,
		blackjack.ThreeName: blackjack.Split,
		blackjack.FourName:  blackjack.Split,
		blackjack.FiveName:  blackjack.Split,
		blackjack.SixName:   blackjack.Split,
		blackjack.SevenName: blackjack.Hit,
		blackjack.EightName: blackjack.Hit,
		blackjack.NineName:  blackjack.Hit,
		blackjack.TenName:   blackjack.Hit,
		blackjack.JackName:  blackjack.Hit,
		blackjack.QueenName: blackjack.Hit,
		blackjack.KingName:  blackjack.Hit,
		blackjack.AceName:   blackjack.Hit,
	},
	// Pair of sevens
	14: {
		blackjack.TwoName:   blackjack.Split,
		blackjack.ThreeName: blackjack.Split,
		blackjack.FourName:  blackjack.Split,
		blackjack.FiveName:  blackjack.Split,
		blackjack.SixName:   blackjack.Split,
		blackjack.SevenName: blackjack.Split,
		blackjack.EightName: blackjack.Hit,
		blackjack.NineName:  blackjack.Hit,
		blackjack.TenName:   blackjack.Hit,
		blackjack.JackName:  blackjack.Hit,
		blackjack.QueenName: blackjack.Hit,
		blackjack.KingName:  blackjack.Hit,
		blackjack.AceName:   blackjack.Hit,
	},
	// Pair of eights
	16: {
		blackjack.TwoName:   blackjack.Split,
		blackjack.ThreeName: blackjack.Split,
		blackjack.FourName:  blackjack.Split,
		blackjack.FiveName:  blackjack.Split,
		blackjack.SixName:   blackjack.Split,
		blackjack.SevenName: blackjack.Split,
		blackjack.EightName: blackjack.Split,
		blackjack.NineName:  blackjack.Split,
		blackjack.TenName:   blackjack.Split,
		blackjack.JackName:  blackjack.Split,
		blackjack.QueenName: blackjack.Split,
		blackjack.KingName:  blackjack.Split,
		blackjack.AceName:   blackjack.Split,
	},
	// Pair of nines
	18: {
		blackjack.TwoName:   blackjack.Split,
		blackjack.ThreeName: blackjack.Split,
		blackjack.FourName:  blackjack.Split,
		blackjack.FiveName:  blackjack.Split,
		blackjack.SixName:   blackjack.Split,
		blackjack.SevenName: blackjack.Stand,
		blackjack.EightName: blackjack.Split,
		blackjack.NineName:  blackjack.Split,
		blackjack.TenName:   blackjack.Stand,
		blackjack.JackName:  blackjack.Stand,
		blackjack.QueenName: blackjack.Stand,
		blackjack.KingName:  blackjack.Stand,
		blackjack.AceName:   blackjack.Stand,
	},
	// Pair of tens
	20: {
		blackjack.TwoName:   blackjack.Stand,
		blackjack.ThreeName: blackjack.Stand,
		blackjack.FourName:  blackjack.Stand,
		blackjack.FiveName:  blackjack.Stand,
		blackjack.SixName:   blackjack.Stand,
		blackjack.SevenName: blackjack.Stand,
		blackjack.EightName: blackjack.Stand,
		blackjack.NineName:  blackjack.Stand,
		blackjack.TenName:   blackjack.Stand,
		blackjack.JackName:  blackjack.Stand,
		blackjack.QueenName: blackjack.Stand,
		blackjack.KingName:  blackjack.Stand,
		blackjack.AceName:   blackjack.Stand,
	},
}