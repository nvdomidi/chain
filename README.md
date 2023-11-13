# Casting Chain

Developing a Chain Casting Tree in Go

# Steps

The following steps are what I think are necessary to develop this project:

## 1. Support Helix

Create a helix (spherical spiral) as demonstrated in ``main.go`` and saved in ``spiral.obj``. You see the helix saved in blender below:

<p align="center">
<img src="pics/helix.png">
</p>

Around each of the points along the helix, a circle must be created. Each circle will be connected to the next circle and triangulated to create a cylinder. Parameters that can be controlled here: 

* Sampling rate along the curve
* Height incrementation between each level
* Radius
* Thickness

## 2. Identifying Pieces

Each piece in the chain must be identified in order. This process includes:

* Separating each piece into its own mesh. This may be done by selecting one triangle and finding all of its neighbors. The code for this is already available. You can also do this in blender by going in edit mode, pressing "P" and clicking "selection by loose parts":

<p align = "center">
<img src="pics/separate.png">
</p>

* Finding the order of the pieces can be done by finding the center point of each separated piece, then somehow finding the orders of those points (maybe form a doubly linked list of closest points?).

## 3. Shape Analysis

This part involves finding the hole. Subtract voxels of shape from its convex hull.