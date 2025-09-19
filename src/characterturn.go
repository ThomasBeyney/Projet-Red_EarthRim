package main

import (
	"fmt"
    "math/rand"
    "time"
)


func CharacterTurn(goblin *Monster, player *Character) bool {
	
	
	for player.Hp > 0 && goblin.Hp > 0 {
    	var choice int
    	fmt.Println("\n=== Interface de combat ===")
		fmt.Println("nom :",player.Name,"  Vie : ",player.Hp,"/",player.MaxHp, "   Mana : ",player.Mana,"/ 50")
    	fmt.Println("1 - Statistique")
		fmt.Println("2 - Attaque Basique ")
   		fmt.Println("3 - Attaque Coup de poing (25 de mana)")
		fmt.Println("4 - Attaque Boule de feu (50 de mana)")
    	fmt.Println("5 - Quitter")
    	fmt.Print("Choix : ")
    	fmt.Scanln(&choice)

		switch choice {
			case 1:
				displayInfo(player)
			case 2:
				fmt.Println("Vous frappez le gobelin avec un coup de basique !")
				goblin.Hp -= 5
				if goblin.Hp < 0 {
					goblin.Hp = 0
				}
				fmt.Println("Il reste", goblin.Hp, "/", goblin.MaxHp, "PV au gobelin.")

				if goblin.Hp == 0 {
					fmt.Println("Le gobelin est vaincu !")
					return false
				}

				goblinPattern(goblin, player)
				if player.Hp == 0 {
					return false
				}

			case 3:
				if player.Mana >= 25 {
					fmt.Println("Vous lancez un Coup de poing sur le gobelin !")
					goblin.Hp -= 8
					player.Mana -= 25
					if goblin.Hp < 0 {
						goblin.Hp = 0
					}
					fmt.Println("Il reste", goblin.Hp, "/", goblin.MaxHp, "PV au gobelin.")
				} else {
					fmt.Println(" ")
					fmt.Println("Pas assez de Mana")
					fmt.Println("Du au manque de mana, votre sort ne ce lance pas, l'ennemie a donc le temps de vous attaquer !")
				}

				if goblin.Hp == 0 {
					fmt.Println("Le gobelin est vaincu !")
					return false
				}

				goblinPattern(goblin, player)
				if player.Hp == 0 {
					return false
				}

			case 4:
				if player.Mana >= 50 {
					fmt.Println("Vous lancez une boule de feu sur le gobelin !")
					goblin.Hp -= 18
					player.Mana -= 50
					if goblin.Hp < 0 {
						goblin.Hp = 0
					}
					fmt.Println("Il reste", goblin.Hp, "/", goblin.MaxHp, "PV au gobelin.")
				} else {
					fmt.Println(" ")
					fmt.Println("Pas assez de Mana")
					fmt.Println("Du au manque de mana, votre sort ne ce lance pas, l'ennemie a donc le temps de vous attaquer !")
				}

				fmt.Println("Il reste", goblin.Hp, "/", goblin.MaxHp, "PV au gobelin.")

				if goblin.Hp == 0 {
					fmt.Println("Le gobelin est vaincu !")
					return false
				}

				goblinPattern(goblin, player)
				if player.Hp == 0 {
					return false
				}

			case 5:
				fmt.Println("Au revoir !")
				return false

			default:
				fmt.Println("Choix invalide, veuillez réessayer.")
			}
		return false
		}
	return false
}





func CombTroll(Troll *Monster, player *Character) bool {
	
	
	for player.Hp > 0 && Troll.Hp > 0 {
    	var choice int
    	fmt.Println("\n=== Interface de combat ===")
		fmt.Println("nom :",player.Name,"  Vie : ",player.Hp,"/",player.MaxHp, "   Mana : ",player.Mana,"/ 50")
    	fmt.Println("1 - Statistique")
		fmt.Println("2 - Attaque Basique ")
   		fmt.Println("3 - Attaque Coup de poing (25 de mana)")
		fmt.Println("4 - Attaque Boule de feu (50 de mana)")
    	fmt.Println("5 - Quitter")
    	fmt.Print("Choix : ")
    	fmt.Scanln(&choice)

		switch choice {
			case 1:
				displayInfo(player)

			case 2:
    			fmt.Println("Vous attaquez le Troll avec un coup basique !")
    			Troll.Hp -= 5
    			if Troll.Hp < 0 {
        			Troll.Hp = 0
    			}
    			fmt.Println("Il reste", Troll.Hp, "/", Troll.MaxHp, "PV au Troll.")

    			if Troll.Hp == 0 {
        			fmt.Println("Le troll est vaincu !")
					player.Exp += 100
					player.Gold += 150
					if player.Exp >= 100 {
						player.Exp = 0
						player.Level += 1
						player.MaxHp += 15
						fmt.Println(" ")
						fmt.Println("Vous gagnez 150 d'or")
						fmt.Println("Bravo ! Vous gagnez un Niveau")
					} else {
						fmt.Println(" ")
						fmt.Println("Vous gagnez 100 points d'expériences et 150 d'or")
					}
        			return false
    			}

				trollPattern(Troll, player)

    			if player.Hp == 0 {
        			return false
    			}
			case 3:
				if player.Mana >= 25 {
					fmt.Println("Vous lancez un Coup de poing sur le troll !")
					Troll.Hp -= 8
					player.Mana -= 25
					if Troll.Hp < 0 {
						Troll.Hp = 0
					}
					fmt.Println("Il reste", Troll.Hp, "/", Troll.MaxHp, "PV au Troll.")
				} else {
					fmt.Println(" ")
					fmt.Println("Pas assez de Mana")
					fmt.Println("Du au manque de mana, votre sort ne ce lance pas, l'ennemie a donc le temps de vous attaquer !")
				}

    			if Troll.Hp == 0 {
        			fmt.Println("Le troll est vaincu !")
					player.Exp += 100
					player.Gold += 150
					if player.Exp >= 100 {
						player.Exp = 0
						player.Level += 1
						player.MaxHp += 15
						fmt.Println(" ")
						fmt.Println("Vous gagnez 150 d'or")
						fmt.Println("Bravo ! Vous gagnez un Niveau")
					} else {
						fmt.Println(" ")
						fmt.Println("Vous gagnez 100 points d'expériences et 150 d'or")
					}
        			return false
    			}

				trollPattern(Troll, player)

    			if player.Hp == 0 {
        			return false
    			}
			case 4:
				if player.Mana >= 50 {
					fmt.Println("Vous lancez une boule de feu sur le troll !")
					Troll.Hp -= 18
					player.Mana -= 50
					if Troll.Hp < 0 {
						Troll.Hp = 0
					}
					fmt.Println("Il reste", Troll.Hp, "/", Troll.MaxHp, "PV au Troll.")
				} else {
					fmt.Println(" ")
					fmt.Println("Pas assez de Mana")
					fmt.Println("Du au manque de mana, votre sort ne ce lance pas, l'ennemie a donc le temps de vous attaquer !")
				}

    			if Troll.Hp == 0 {
        			fmt.Println("Le troll est vaincu !")
					player.Exp += 100
					player.Gold += 150
					if player.Exp >= 100 {
						player.Exp = 0
						player.Level += 1
						player.MaxHp += 15
						fmt.Println(" ")
						fmt.Println("Vous gagnez 150 d'or")
						fmt.Println("Bravo ! Vous gagnez un Niveau")
					} else {
						fmt.Println(" ")
						fmt.Println("Vous gagnez 100 points d'expériences et 150 d'or")
					}
        			return false
    			}

				trollPattern(Troll, player)

    			if player.Hp == 0 {
        			return false
    			}
			case 5:
				fmt.Println("Au revoir !")
				return false

			default:
				fmt.Println("Choix invalide, veuillez réessayer.")
			}
		return false
		}
	return false
}





func CombLoup(Loup *Monster, player *Character) bool {
	
	
	for player.Hp > 0 && Loup.Hp > 0 {
    	var choice int
    	fmt.Println("\n=== Interface de combat ===")
		fmt.Println("nom :",player.Name,"  Vie : ",player.Hp,"/",player.MaxHp, "   Mana : ",player.Mana,"/ 50")
    	fmt.Println("1 - Statistique")
		fmt.Println("2 - Attaque Basique ")
   		fmt.Println("3 - Attaque Coup de poing (25 de mana)")
		fmt.Println("4 - Attaque Boule de feu (50 de mana)")
    	fmt.Println("5 - Quitter")
    	fmt.Print("Choix : ")
    	fmt.Scanln(&choice)

		switch choice {
			case 1:
				displayInfo(player)

			case 2:
    			fmt.Println("Vous attaquez le Loup avec un coup basique !")
    			Loup.Hp -= 5
    			if Loup.Hp < 0 {
        			Loup.Hp = 0
    			}
    			fmt.Println("Il reste", Loup.Hp, "/", Loup.MaxHp, "PV au Loup.")

    			if Loup.Hp == 0 {
        			fmt.Println("Le loup est vaincu !")
					rand.Seed(time.Now().UnixNano())
					n := rand.Intn(25)
					player.Exp += 15
					player.Gold += n
					if player.Exp >= 100 {
						player.Exp = 0
						player.Level += 1
						player.MaxHp += 15
						fmt.Println(" ")
						fmt.Println("Vous gagnez",n,"d'or")
						fmt.Println("Bravo ! Vous gagnez un Niveau")
					} else {
						fmt.Println(" ")
						fmt.Println("Vous gagnez 15 points d'expériences et",n,"d'or")
					}
        			return false
    			}

				loupPattern(Loup, player)

    			if player.Hp == 0 {
        			return false
    			}
			case 3:
    			if player.Mana >= 25 {
					fmt.Println("Vous lancez un Coup de poing sur le loup !")
					Loup.Hp -= 8
					player.Mana -= 25
					if Loup.Hp < 0 {
						Loup.Hp = 0
					}
					fmt.Println("Il reste", Loup.Hp, "/", Loup.MaxHp, "PV au loup.")
				} else {
					fmt.Println(" ")
					fmt.Println("Pas assez de Mana")
					fmt.Println("Du au manque de mana, votre sort ne ce lance pas, l'ennemie a donc le temps de vous attaquer !")
				}

    			if Loup.Hp == 0 {
        			fmt.Println("Le loup est vaincu !")
					rand.Seed(time.Now().UnixNano())
					n := rand.Intn(25)
					player.Exp += 15
					player.Gold += n
					if player.Exp >= 100 {
						player.Exp = 0
						player.Level += 1
						player.MaxHp += 15
						fmt.Println(" ")
						fmt.Println("Vous gagnez",n,"d'or")
						fmt.Println("Bravo ! Vous gagnez un Niveau")
					} else {
						fmt.Println(" ")
						fmt.Println("Vous gagnez 15 points d'expériences et",n,"d'or")
					}
        			return false
    			}

				loupPattern(Loup, player)

    			if player.Hp == 0 {
        			return false
    			}
			case 4:
    			if player.Mana >= 50 {
					fmt.Println("Vous lancez une boule de feu sur le loup !")
					Loup.Hp -= 18
					player.Mana -= 50
					if Loup.Hp < 0 {
						Loup.Hp = 0
					}
					fmt.Println("Il reste", Loup.Hp, "/", Loup.MaxHp, "PV au Troll.")
				} else {
					fmt.Println(" ")
					fmt.Println("Pas assez de Mana")
					fmt.Println("Du au manque de mana, votre sort ne ce lance pas, l'ennemie a donc le temps de vous attaquer !")
				}

    			if Loup.Hp == 0 {
        			fmt.Println("Le loup est vaincu !")
					rand.Seed(time.Now().UnixNano())
					n := rand.Intn(25)
					player.Exp += 15
					player.Gold += n
					if player.Exp >= 100 {
						player.Exp = 0
						player.Level += 1
						player.MaxHp += 15
						fmt.Println(" ")
						fmt.Println("Vous gagnez",n,"d'or")
						fmt.Println("Bravo ! Vous gagnez un Niveau")
					} else {
						fmt.Println(" ")
						fmt.Println("Vous gagnez 15 points d'expériences et",n,"d'or")
					}
        			return false
    			}

				loupPattern(Loup, player)

    			if player.Hp == 0 {
        			return false
    			}
			case 5:
				fmt.Println("Au revoir !")
				return false

			default:
				fmt.Println("Choix invalide, veuillez réessayer.")
			}
		return false
		}
	return false
}




func CombOrc(Orc *Monster, player *Character) bool {
	
	
	for player.Hp > 0 && Orc.Hp > 0 {
    	var choice int
    	fmt.Println("\n=== Interface de combat ===")
		fmt.Println("nom :",player.Name,"  Vie : ",player.Hp,"/",player.MaxHp, "   Mana : ",player.Mana,"/ 50")
    	fmt.Println("1 - Statistique")
		fmt.Println("2 - Attaque Basique ")
   		fmt.Println("3 - Attaque Coup de poing (25 de mana)")
		fmt.Println("4 - Attaque Boule de feu (50 de mana)")
    	fmt.Println("5 - Quitter")
    	fmt.Print("Choix : ")
    	fmt.Scanln(&choice)

		switch choice {
			case 1:
				displayInfo(player)

			case 2:
    			fmt.Println("Vous attaquez l'Orc avec une attaque basique !")
    			Orc.Hp -= 5
    			if Orc.Hp < 0 {
        			Orc.Hp = 0
    			}
    			fmt.Println("Il reste", Orc.Hp, "/", Orc.MaxHp, "PV au Orc.")

    			if Orc.Hp == 0 {
        			fmt.Println("L'Orc est vaincu !")
					rand.Seed(time.Now().UnixNano())
					n := rand.Intn(60)
					player.Exp += 45
					player.Gold += n
					if player.Exp >= 100 {
						player.Exp = 0
						player.Level += 1
						player.MaxHp += 15
						fmt.Println(" ")
						fmt.Println("Vous gagnez",n,"d'or")
						fmt.Println("Bravo ! Vous gagnez un Niveau")
					} else {
						fmt.Println(" ")
						fmt.Println("Vous gagnez 45 points d'expériences et",n,"d'or")
					}
        			return false
    			}

				orcPattern(Orc, player)

    			if player.Hp == 0 {
        			return false
    			}

			case 3:
    			if player.Mana >= 25 {
					fmt.Println("Vous lancez un Coup de poing sur le orc !")
					Orc.Hp -= 8
					player.Mana -= 25
					if Orc.Hp < 0 {
						Orc.Hp = 0
					}
					fmt.Println("Il reste", Orc.Hp, "/", Orc.MaxHp, "PV au gobelin.")
				} else {
					fmt.Println(" ")
					fmt.Println("Pas assez de Mana")
					fmt.Println("Du au manque de mana, votre sort ne ce lance pas, l'ennemie a donc le temps de vous attaquer !")
				}

    			if Orc.Hp == 0 {
        			fmt.Println("L'Orc est vaincu !")
					rand.Seed(time.Now().UnixNano())
					n := rand.Intn(60)
					player.Exp += 45
					player.Gold += n
					if player.Exp >= 100 {
						player.Exp = 0
						player.Level += 1
						player.MaxHp += 15
						fmt.Println(" ")
						fmt.Println("Vous gagnez",n,"d'or")
						fmt.Println("Bravo ! Vous gagnez un Niveau")
					} else {
						fmt.Println(" ")
						fmt.Println("Vous gagnez 45 points d'expériences et",n,"d'or")
					}
        			return false
    			}

				orcPattern(Orc, player)

    			if player.Hp == 0 {
        			return false
    			}
			case 4:
    			if player.Mana >= 50 {
					fmt.Println("Vous lancez une boule de feu sur le orc !")
					Orc.Hp -= 18
					player.Mana -= 50
					if Orc.Hp < 0 {
						Orc.Hp = 0
					}
					fmt.Println("Il reste", Orc.Hp, "/", Orc.MaxHp, "PV au Troll.")
				} else {
					fmt.Println(" ")
					fmt.Println("Pas assez de Mana")
					fmt.Println("Du au manque de mana, votre sort ne ce lance pas, l'ennemie a donc le temps de vous attaquer !")
				}

    			if Orc.Hp == 0 {
        			fmt.Println("L'Orc est vaincu !")
					rand.Seed(time.Now().UnixNano())
					n := rand.Intn(60)
					player.Exp += 45
					player.Gold += n
					if player.Exp >= 100 {
						player.Exp = 0
						player.Level += 1
						player.MaxHp += 15
						fmt.Println(" ")
						fmt.Println("Vous gagnez",n,"d'or")
						fmt.Println("Bravo ! Vous gagnez un Niveau")
					} else {
						fmt.Println(" ")
						fmt.Println("Vous gagnez 45 points d'expériences et",n,"d'or")
					}
        			return false
    			}

				orcPattern(Orc, player)

    			if player.Hp == 0 {
        			return false
    			}

			case 5:
				fmt.Println("Au revoir !")
				return false

			default:
				fmt.Println("Choix invalide, veuillez réessayer.")
			}
		return false
		}
	return false
}
