# Go Business Management

[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE)

This is a port of Go of a BASIC strategy/management game called Business Management. he BASIC code was originally published in 1979 within [Stimulating Simulations 2nd Revised edition by Engel, C.W](https://bookwyrm.social/book/50763/s/stimulating-simulations).

In this simulation you manage a small factory that produces three different kinds of products (P1 - P3). There are three different raw materials with each product requiring two materials in order to be manufactured. For example in order to manufacture a unit of P1 you would need one unit each of R2 and R3, to manufacture a unit of P2 you would need one unit each of R1 and R3.

Raw material cost will vary from $10 to $20 per unit while it costs between $1 to $9 per unit to manufacture a product. Similarly, the selling price of each product will vary from $50 to $90 per unit.

The simulation runs for twelve months, after which all stock on hand is sold and your final profit (score) is calculated. Each month you're presented with the current material costs and product price and asked if you want to buy materials, manufacture or sell products. You can only do one of these actions per month.

Build using `go build` and then execute, you will be presented with output like so:

```
Item:  Materials: Product:
1      0 $14      0 $77
2      0 $16      0 $69
3      0 $16      0 $54
Month 1, you have $500
Manufacturing costs are $2/unit
Transaction (O,B,M,S) ? 
```