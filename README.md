# Pizza Delivery Challenge

## Overview

After reading the problem, my first instinct is to define/model the problem scope so
that it's clear to me what the rules are and this will help define a solution:

### Actors
 - `delivery person`
 - `goat`

### Actions
- `move` - an actor can move on his turn
- `stay` - only applies to the first action when a pizza is initially delivered before the first `move`
- `deliver` - delivers a pizza

The delivery person always makes the first move. An actor can take the following actions which may be considered atomic:
```
(move|stay) + deliver
```
Either `stay` and `deliver` a pizza OR `move` and `deliver` a pizza

### Context

The actions defined by the dispatcher are relative to the current location of an actor. We need a way to define an absolute location independent of the actors' positions. The challenge defines a two-dimensional grid as the landscape where the challenge takes place. We can model this landscape as a Cartesian plane where houses are synonymous with points on the plane and are identified by their positions along the `x-axis` and `y-axis`.

### Putting it all together

Now that I have an idea of how it all fits together, let's define how moves will be traced along the Cartesian plane. An actor can move in any of the following directions: `^` `v` `>` `<`

Each direction represents a cardinal point: `north` `south` `east` `west`. We can model this on the plane where `north` and `south` are measured along the `y-axis` and `east` and `west` are measured along the `x-axis`.

So, assuming the challenge starts at `origin` `(0,0)` (and a pizza is delivered there), then `^` would indicate "move one position `north` and deliver a pizza". The actor would then be at `(0,1)`.
