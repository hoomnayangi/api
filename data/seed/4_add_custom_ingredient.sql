BEGIN TRANSACTION;

ALTER TABLE ingredients 
ADD COLUMN is_hidden BOOLEAN DEFAULT false;

UPDATE ingredients
SET is_hidden = true
WHERE name LIKE '%gà%';

UPDATE ingredients
SET is_hidden = true
WHERE name LIKE '%bò%';

UPDATE ingredients
SET is_hidden = true
WHERE name LIKE '%heo%';

INSERT INTO ingredients(name, id) VALUES ('gà', 1100);
INSERT INTO ingredients(name, id) VALUES ('bò', 1101);
INSERT INTO ingredients(name, id) VALUES ('heo', 1102);

ALTER SEQUENCE public.ingredients_id_seq RESTART WITH 1279;


INSERT INTO recipes_ingredients(recipe_id, ingredient_id, amount) 
(SELECT DISTINCT(recipe_id), 1100, 1  from recipes_ingredients INNER JOIN ingredients on ingredients.id = recipes_ingredients.ingredient_id AND name LIKE '%gà%');

INSERT INTO recipes_ingredients(recipe_id, ingredient_id, amount) 
(SELECT DISTINCT(recipe_id), 1101, 1  from recipes_ingredients INNER JOIN ingredients on ingredients.id = recipes_ingredients.ingredient_id AND name LIKE '%bò%');

INSERT INTO recipes_ingredients(recipe_id, ingredient_id, amount) 
(SELECT DISTINCT(recipe_id), 1102, 1  from recipes_ingredients INNER JOIN ingredients on ingredients.id = recipes_ingredients.ingredient_id AND name LIKE '%heo%');
COMMIT;