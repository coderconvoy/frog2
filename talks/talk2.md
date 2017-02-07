Well Greetings Webbed Folk.

Today we are continuing working out how to use Engo.io, the 2D game engine from Paked (his git hub handle) et al.

The aim of this tutorial is to make a game something like the classic frogger game.

So far we have something like this:

I've been working ahead, and I have made the next few steps too, and that now looks like this:

The plan today is to add cars, and give them movement.

To do this I will create a car motion system, which loops through all the cars, and moves them according to their velocity.
I will also need to create a car type with the componenets needed for this system including common.SpaceComponent from Engo, and our own VelocityComponent.

First, let's look at all the code we have so far, to make some sense of where we are going.


