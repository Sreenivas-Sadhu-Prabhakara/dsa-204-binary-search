# Binary Search

> NeetCode 150 · Binary Search — the plain closed-interval form, written once, carefully: inclusive bounds, `lo <= hi`, overflow-safe midpoint. Go, stdlib only, tested.

## Problem

Index of a target in a sorted array, or -1. The primitive itself — which this series uses dozens of times dressed up (rotated arrays, answer spaces, matrices) — here in its unadorned form, because the dressed-up versions fail exactly where the plain one was fuzzy.

## Approach

The closed-interval contract: `[lo, hi]` always contains the target if it exists; `lo <= hi` keeps searching while the interval is non-empty; a miss shrinks the interval past the probe (`mid ± 1`, legal only because bounds are inclusive). The midpoint `lo + (hi-lo)/2` avoids the signed-overflow bug that lived in production libraries for decades (JDK bug 5045582). Jon Bentley's finding stands: most professionals get this wrong on paper — the three decisions (inclusive vs exclusive bounds, loop condition, shrink step) must agree, and each half-open variant is fine *if consistent*. In real Go, `sort.SearchInts` is the tool; writing it raw is the exercise.

**Complexity:** O(log n) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
