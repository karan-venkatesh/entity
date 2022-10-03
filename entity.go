package entity

import (
    "math/rand"
    "time"
    "fmt"
)

type EntitySpawn struct{
    size int
    speed int
    attack int
    defence int
    herb bool
    carn bool 
}


// init sets initial values for variables used in the function.
func init() {
    rand.Seed(time.Now().UnixNano())
}

func GetEntityDescription(en EntitySpawn) string{
    return fmt.Sprintf("Size:\t\t%d\nSpd:\t\t%d\nAtk:\t\t%d\nDef:\t\t%d\nHerb:\t\t%t\nCarn:\t\t%t\n",
        en.size,
        en.speed,
        en.attack,
        en.defence,
        en.herb,
        en.carn, 
    )
}

func computeEntitySize(en EntitySpawn) int {
    var herb_score int
    if en.herb {
        herb_score = 4
    }

    var carn_score int
    if en.carn {
        carn_score = 6
    }

    return en.speed + en.attack + en.defence + herb_score + carn_score
}

func GenerateRandomEntity() EntitySpawn {
    spawned_entity := EntitySpawn{
        speed: rand.Intn(10) + 1,
        attack: rand.Intn(10) + 1,
        defence: rand.Intn(10) + 1,
        herb: rand.Intn(2) == 1,
        carn: rand.Intn(2) == 1,
    }

    spawned_entity.size = computeEntitySize(spawned_entity)
	return spawned_entity
}

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func MutateEntity(parent EntitySpawn, mutation_rate int) EntitySpawn {
    herb := parent.herb
    if rand.Intn(100 - mutation_rate) == 1 {
        herb = !herb
    }

    carn := parent.carn
    if rand.Intn(100 - mutation_rate) == 1 {
        carn = !carn 
    }

    child := EntitySpawn{
        speed: max(parent.speed + rand.Intn(mutation_rate) - mutation_rate / 2, 1),
        attack: max(parent.attack + rand.Intn(mutation_rate) - mutation_rate / 2, 1),
        defence: max(parent.defence + rand.Intn(mutation_rate) - mutation_rate / 2, 1),
        herb: herb,
        carn: carn,
    }
    child.size = computeEntitySize(child)
    return child
}


func Fight(entity_a EntitySpawn, entity_b EntitySpawn) (int, int) {
    var entity_a_damage int
    var entity_b_damage int

    entity_a_damage = entity_b.attack / entity_a.defence 
    entity_b_damage = 2 * entity_a.attack / entity_b.defence

    return entity_a_damage, entity_b_damage
}

