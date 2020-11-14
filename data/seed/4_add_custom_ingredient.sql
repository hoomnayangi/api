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
(SELECT DISTINCT(recipe_id), 1100, 0  from recipes_ingredients INNER JOIN ingredients on ingredients.id = recipes_ingredients.ingredient_id AND (lower(name) LIKE '%gà%' and lower(name) not like '%trứng gà%') );

INSERT INTO recipes_ingredients(recipe_id, ingredient_id, amount) 
SELECT DISTINCT(recipe_id), (SELECT id from ingredients where lower(name) = 'trứng gà'), 1  from recipes_ingredients INNER JOIN ingredients on ingredients.id = recipes_ingredients.ingredient_id AND lower(name) LIKE '%trứng gà%' on conflict do nothing;

INSERT INTO recipes_ingredients(recipe_id, ingredient_id, amount) 
(SELECT DISTINCT(recipe_id), 1101, 0  from recipes_ingredients INNER JOIN ingredients on ingredients.id = recipes_ingredients.ingredient_id AND lower(name) LIKE '%bò%');

INSERT INTO recipes_ingredients(recipe_id, ingredient_id, amount) 
(SELECT DISTINCT(recipe_id), 1102, 0  from recipes_ingredients INNER JOIN ingredients on ingredients.id = recipes_ingredients.ingredient_id AND lower(name) LIKE '%heo%');
COMMIT;

UPDATE ingredients
SET picture = 'https://i.imgur.com/6tVrFYf.png'
WHERE LOWER(name) = 'bò';


UPDATE ingredients
SET picture = 'https://i.imgur.com/k8CEcbp.png'
WHERE LOWER(name) = 'gà';


UPDATE ingredients
SET picture = 'https://i.imgur.com/MFY8jTD.png'
WHERE LOWER(name) = 'heo';


UPDATE ingredients
SET picture = 'https://i.imgur.com/TyP2209.png'
WHERE LOWER(name) = 'tỏi'; 


UPDATE ingredients
SET picture = 'https://i.imgur.com/hLcKT3D.png'
WHERE LOWER(name) = 'cá'; 



UPDATE ingredients
SET picture = 'https://i.imgur.com/Rb9YcpW.png'
WHERE LOWER(name) = 'cá ngừ'; 


UPDATE ingredients
SET picture = 'https://i.imgur.com/RHdaJcU.png'
WHERE LOWER(name) = 'cá lăng'; 


UPDATE ingredients
SET picture = 'https://i.imgur.com/sU7Sfms.png'
WHERE LOWER(name) = 'ớt';

UPDATE ingredients
SET picture = 'https://png.pngtree.com/png-clipart/20190929/ourlarge/pngtree-illustration-chicken-egg-vector-png-image_1747831.jpg'
WHERE LOWER(name) = 'trứng gà';