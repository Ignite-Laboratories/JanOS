// Package mirror is used to generate mirror types of another structure.  This is useful when you have a type
// which takes in many generics for flexibility, but want to expose a more generalized variant of the same type.
// Source and target types are selected before the methods of the source type are mirrored as passthrough from
// the target.
package mirror
