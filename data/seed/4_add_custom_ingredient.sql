BEGIN TRANSACTION;

ALTER TABLE ingredients 
ADD COLUMN is_hidden BOOLEAN DEFAULT false;

UPDATE ingredients
SET is_hidden = true
WHERE lower(name) LIKE lower('%gà%') and lower(name) <> 'trứng gà';

UPDATE ingredients
SET is_hidden = true
WHERE lower(name) LIKE lower('%bò%');

UPDATE ingredients
SET is_hidden = true
WHERE lower(name) LIKE lower('%heo%');

INSERT INTO ingredients(name, id) VALUES ('Gà', 1100);
INSERT INTO ingredients(name, id) VALUES ('Bò', 1101);
INSERT INTO ingredients(name, id) VALUES ('Heo', 1102);

ALTER SEQUENCE public.ingredients_id_seq RESTART WITH 1279;


INSERT INTO recipes_ingredients(recipe_id, ingredient_id, amount) 
(SELECT DISTINCT(recipe_id), 1100, 0  from recipes_ingredients INNER JOIN ingredients on ingredients.id = recipes_ingredients.ingredient_id AND lower(name) LIKE '%gà%');

INSERT INTO recipes_ingredients(recipe_id, ingredient_id, amount) 
SELECT DISTINCT(recipe_id), (SELECT id from ingredients where lower(name) = 'trứng gà'), 1  from recipes_ingredients INNER JOIN ingredients on ingredients.id = recipes_ingredients.ingredient_id AND lower(name) LIKE '%trứng gà%' on conflict do nothing;

INSERT INTO recipes_ingredients(recipe_id, ingredient_id, amount) 
(SELECT DISTINCT(recipe_id), 1101, 0  from recipes_ingredients INNER JOIN ingredients on ingredients.id = recipes_ingredients.ingredient_id AND lower(name) LIKE '%bò%');

INSERT INTO recipes_ingredients(recipe_id, ingredient_id, amount) 
(SELECT DISTINCT(recipe_id), 1102, 0  from recipes_ingredients INNER JOIN ingredients on ingredients.id = recipes_ingredients.ingredient_id AND lower(name) LIKE '%heo%');
COMMIT;