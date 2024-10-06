While you could use DISTINCT with SUM or AVG, in practice it's rare to want to just sum or average just the unique values. 
When it comes to MAX and MIN, they aren't affected by DISTINCT – whether there are duplicates or not, 
the lowest/highest value in the dataset will be the same.

---
You can use DISTINCT with aggregate functions – the most common one being COUNT. Here's an example that finds the number of unique user's who made trades:

SELECT COUNT(DISTINCT user_id) 
FROM trades;

---

Assume you're given a table containing data on Amazon customers and their spending on products in different category. Write a query using COUNT DISTINCT to identify the number of unique products within each product category.

SELECT category, COUNT(DISTINCT product)
FROM product_spend
GROUP BY category;