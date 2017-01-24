Greetings, People of the Web.

Today we are looking at the "Engo, 2D Engine" for go lang.  It's a pretty sweet engine, for getting 2d games running. Sadly right now, it's not had a lot of activity for about half a year, so this video, is part of my hopes re-invigorate it.

A quick overview of the package. Engo is an engine, build on top of Open GL using the Entity, Component, System architecture (Or ECS). Once you get your head around ECS, this whole thing starts to come together.

You can download it from engo.io (that's the website) or from github. It has a few dependencies, which are mentioned in the readme.md file so make sure you have those before running it. (I didn't see that, and it took me a while to find all the things I needed so it took me a while).

ECS
====

Entities
---------
The Entity component system, is a relatively flexible way of working. I'll try to give you my overview:

An Entity, is the equivilent of an Object on OOP each item in your world will be an Entity. Entities hold the data, but do not manipulate it. Which is the biggest difference between this and OOP.

Components
----------
Each entity is the combination of a number of Components. These are normally included in the Entity type as parts annonymous struct members.

eg :
    type FrogEntity struct{
        *ecs.BasicEntity //(Really a component, but provides a UniqueID)
        *common.SpaceComponent // common is a library of useful components and systems
        *common.RenderComponent
        *JumpComponent // my own component, which I want to keep track of height or something like it.
    }


Lastly Systems.
-----------

Systems, actually do the things. Each system, can have objects added to it, and every update cycle will basically loop through all of the items and do whatever it wants to them.

So the rendering system keeps a list of []struct{BasicEntity,SpaceComponent,RenderComponent} and draws them on every update cycle.

I might make a JumpSystem, which stores a []struct{BasicEntity,SpaceComponent,JumpComponent}

And every update, it will make note of every item that can jump, and move it like that.

in Engo, System is an interface with the following 2 methods 
*   Update(d float32) //Do the things
*   Remove(e ecs.BasicEntity)// Find element with same ID as 'e' and remove it

Most of them also have an Add method, which accepts each of the components needed. 
They can also have a New method, which allows them to get started with access to the world they are part of.

func (rs *RenderSystem) Add(eb *ecs.BasicEntity , sc *SpaceComponent, rc *RenderComponent){
    rs.obs = append(rs.obs , RenderEntity{eb,sc,rc}
}

World
-----

In Engo, World is the system that tells all the systems to update. Once you add a system to the world.

And Scene, as a way of breaking up views and more separate parts of the project.

Hopefully this will become clearer as we get coding.








