package server

import "github.com/google/uuid"

var Jeoqcs = [5]QuestionCategory{
	{
		Name: "A WRITER'S LIFE FOR ME",
		Qs: [5]Question{
			{
				Id:       uuid.New(),
				Q:        "Corvus corax is the scientific name of this crow cousin",
				A:        "A raven",
				V:        200,
				Revealed: false,
			},
			{
				Id:       uuid.New(),
				Q:        "5-letter magazine founded in 1945 by John H. Johnson",
				A:        "Ebony",
				V:        100,
				Revealed: false,
			},
			{
				Id:       uuid.New(),
				Q:        "Ancient Egyptians used this 4-letter cosmetic to line the eyes",
				A:        "Kohl",
				V:        200,
				Revealed: false,
			},
			{
				Id:       uuid.New(),
				Q:        "The RAF's Gloster Meteor, for example",
				A:        "A jet",
				V:        100,
				Revealed: false,
			},
			{
				Id:       uuid.New(),
				Q:        "Residue of coal tar distillation used to pave roads & waterproof roofs",
				A:        "Pitch",
				V:        200,
				Revealed: false,
			},
		},
	},
	{
		Name: "SYNONYMS FOR BLACK",
		Qs: [5]Question{
			{
				Id:       uuid.New(),
				Q:        "Corvus corax is the scientific name of this crow cousin",
				A:        "A raven",
				V:        100,
				Revealed: false,
			},
			{
				Id:       uuid.New(),
				Q:        "5-letter magazine founded in 1945 by John H. Johnson",
				A:        "Ebony",
				V:        300,
				Revealed: false,
			},
			{
				Id:       uuid.New(),
				Q:        "Ancient Egyptians used this 4-letter cosmetic to line the eyes",
				A:        "Kohl",
				V:        100,
				Revealed: false,
			},
			{
				Id:       uuid.New(),
				Q:        "The RAF's Gloster Meteor, for example",
				A:        "A jet",
				V:        300,
				Revealed: false,
			},
			{
				Id:       uuid.New(),
				Q:        "Residue of coal tar distillation used to pave roads & waterproof roofs",
				A:        "Pitch",
				V:        100,
				Revealed: false,
			},
		},
	},
	{
		Name: "THE NEW NEWS",
		Qs: [5]Question{
			{
				Id:       uuid.New(),
				Q:        "It's inserted into the crankcase to check a car's oil level",
				A:        "A dipstick",
				V:        200,
				Revealed: false,
			},
			{
				Id:       uuid.New(),
				Q:        "It's the term for weight placed low in a ship to make it more stable & seaworthy",
				A:        "Ballast",
				V:        100,
				Revealed: false,
			},
			{
				Id:       uuid.New(),
				Q:        "Landing a plane in this, an airflow that hits broadside, can be challenging as seen here",
				A:        "Crosswind",
				V:        200,
				Revealed: false,
			},
			{
				Id:       uuid.New(),
				Q:        "A horse rented out for riding is known as this, also a slang term for a taxi driver",
				A:        "Hack",
				V:        100,
				Revealed: false,
			},
			{
				Id:       uuid.New(),
				Q:        "This 7-letter activity is the K in the WKA, an association excited about going 100 mph, 1 inch off the ground",
				A:        "Karting",
				V:        200,
				Revealed: false,
			},
		},
	},
	{
		Name: "TRANSPORTATION TERMS",
		Qs: [5]Question{
			{
				Id:       uuid.New(),
				Q:        "Corvus corax is the scientific name of this crow cousin",
				A:        "A raven",
				V:        100,
				Revealed: false,
			},
			{
				Id:       uuid.New(),
				Q:        "5-letter magazine founded in 1945 by John H. Johnson",
				A:        "Ebony",
				V:        300,
				Revealed: false,
			},
			{
				Id:       uuid.New(),
				Q:        "Ancient Egyptians used this 4-letter cosmetic to line the eyes",
				A:        "Kohl",
				V:        100,
				Revealed: false,
			},
			{
				Id:       uuid.New(),
				Q:        "The RAF's Gloster Meteor, for example",
				A:        "A jet",
				V:        300,
				Revealed: false,
			},
			{
				Id:       uuid.New(),
				Q:        "Residue of coal tar distillation used to pave roads & waterproof roofs",
				A:        "Pitch",
				V:        100,
				Revealed: false,
			},
		},
	},
	{
		Name: "British History",
		Qs: [5]Question{
			{
				Id:       uuid.New(),
				Q:        "It's inserted into the crankcase to check a car's oil level",
				A:        "A dipstick",
				V:        200,
				Revealed: false,
			},
			{
				Id:       uuid.New(),
				Q:        "It's the term for weight placed low in a ship to make it more stable & seaworthy",
				A:        "Ballast",
				V:        100,
				Revealed: false,
			},
			{
				Id:       uuid.New(),
				Q:        "Landing a plane in this, an airflow that hits broadside, can be challenging as seen here",
				A:        "Crosswind",
				V:        200,
				Revealed: false,
			},
			{
				Id:       uuid.New(),
				Q:        "A horse rented out for riding is known as this, also a slang term for a taxi driver",
				A:        "Hack",
				V:        100,
				Revealed: false,
			},
			{
				Id:       uuid.New(),
				Q:        "This 7-letter activity is the K in the WKA, an association excited about going 100 mph, 1 inch off the ground",
				A:        "Karting",
				V:        200,
				Revealed: false,
			},
		},
	},
}
