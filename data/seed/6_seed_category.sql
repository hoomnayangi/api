ALTER TABLE recipes
ADD COLUMN category varchar(20);

UPDATE recipes
SET category=subquery.category from 
(select id, CASE WHEN id%3 = 2 THEN '2'
         WHEN id%3 = 1 THEN '1'
         ELSE '3' END
         AS category
from recipes) as subquery where recipes.id = subquery.id;
