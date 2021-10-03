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

## Considerations

It doesn't sound like the dispatcher is trustworthy so I opted to read and validate the entire input at once before continuing to process it. This means that I'm reading the full string of data into memory. This isn't an optimal solution as the amount of memory we need will grow as the length of the string grows. A better solution would be to read the string a few (debatable how many) characters at a time and process the data in batches. The concern here is that if there is an invalid instructioni/character in the input as we process it in batches, what is the appropriate response? Should that deliverer just skip it or does this mean the remaining input is also invalid?

## Answers

### Part 1

Only the delivery person is available.

The number of houses that receive at least 1 pizza: `2565`

```bash
go run main.go -file-input directions/pizza_delivery_input.txt -deliverer-count 1
```

### Part 2

The delivery person and goat are available.

The number of houses that receive at least 1 pizza: `2639`

```bash
go run main.go -file-input directions/pizza_delivery_input.txt -deliverer-count 2
```
