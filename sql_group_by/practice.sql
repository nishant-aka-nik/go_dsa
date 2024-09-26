-- here no need to give ticker in column it will still group by the ticker column
SELECT MIN(low) as price from stock_prices 
GROUP BY ticker 
ORDER BY price DESC


-- error happened here 
-- column "stock_prices.low" must appear in the GROUP BY clause or be used in an aggregate function (LINE: 3)
-- as when we use the min the column name will be min and when the order by tries  to order the data it wont find the column as min 
-- it will throw error 
SELECT MIN(low) from stock_prices 
GROUP BY ticker 
ORDER BY low DESC
-- fixed query -- we will use the alias and then order by will order by prices  
SELECT ticker, MIN(low) as prices from stock_prices 
GROUP BY ticker 
ORDER BY prices DESC

--
--The GROUP BY clause is used in conjunction with aggregate functions to group rows 
--that have the same values in specified columns into summary rows. 
--It allows you to perform aggregate calculations on each group independently.

SELECT column1, AGGREGATE_FUNCTION(column2)
FROM table_name
WHERE condition
GROUP BY column1;

--------------------------------------------------------------------------------------------------------
-- group by with aggregate functions chat gpt examples
-- https://chatgpt.com/share/66f596dc-c598-8006-b4d2-05ea13bc9bfd


--------------------------------------------------------------------------------------------------------
Q. How many candidates possess each of the different skills?
Ans. How many candidate ->>> matlab count lagega candidate column pe 

skill ke basis pe group kiya or uss group pe count aggregate kiya

SELECT skill, COUNT(candidate_id)
FROM candidates
GROUP BY skill
ORDER BY count DESC

--------------------------------------------------------------------------------------------------------
it's absolutely possible for a query to have both GROUP BY and ORDER BY, with GROUP BY coming first!
--------------------------------------------------------------------------------------------------------


