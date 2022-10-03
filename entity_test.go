package entity

import (
    "testing"
	"fmt"
	"math/rand"

)




func TestGenerateEntity(t *testing.T) {

    entity_spawned := GenerateRandomEntity()
	fmt.Println(GetEntityDescription(entity_spawned))

    // if entity_name.carn {
    //     t.Fatalf(`Entity name is too short`)
    // }
}

func TestGenerateChildEntity(t *testing.T) {

    entity_spawned := GenerateRandomEntity()
	entity_spawned_child := MutateEntity(entity_spawned,5)
	fmt.Println(GetEntityDescription(entity_spawned))
	fmt.Println(GetEntityDescription(entity_spawned_child))

    // if entity_name.carn {
    //     t.Fatalf(`Entity name is too short`)
    // }
}


func TestFightTwoEntities(t *testing.T) {

    entity_a := GenerateRandomEntity()
	entity_b := GenerateRandomEntity()
	fmt.Println(GetEntityDescription(entity_a))
	fmt.Println(GetEntityDescription(entity_b))
	fmt.Println(Fight(entity_a,entity_b))

    // if entity_name.carn {
    //     t.Fatalf(`Entity name is too short`)
    // }
}


func TestDaySimulation(t *testing.T) {
	var creatures []EntitySpawn

	var hp []int
	for i :=0; i < 5; i++ {
		entity_new := GenerateRandomEntity()
		creatures = append(creatures, entity_new)
		hp = append(hp,entity_new.size)
		fmt.Println(GetEntityDescription(creatures[i]))
	}

	for i := 1; i<20; i++{
		for j, c := range creatures {
			if i % c.speed == 0 {
				fmt.Println(i,j,GetEntityDescription(c))
				target_index := rand.Intn(len(creatures))
				damage_a, damage_b := Fight(c,creatures[target_index])
				hp[j] -= damage_a
				hp[target_index] -= damage_b
			}
		}
	}

	fmt.Println(hp)
}