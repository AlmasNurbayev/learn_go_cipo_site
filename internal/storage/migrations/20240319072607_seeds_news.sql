-- +goose Up
-- +goose StatementBegin
IF EXISTS (
  SELECT
    1
  FROM
    your_table
) BEGIN -- Здесь можно указать действия, которые нужно выполнить,
-- если записи в таблице есть
PRINT 'Записи в таблице существуют';

END -- добавляем новые записи
ELSE BEGIN
INSERT INTO
  news (title, data, image_path)
VALUES
  (
    'У нас новые классные сандалии!',
    '<p>Все как обычно - натуральная кожа, отличное качество, стильный дизайн</p>
<p>С 26 апреля в продаже (наличие и цены - картинки кликабельны), размеры 21-36:</p>
<div>
<a href="/catalog/12355/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/6163-837-01.jpg"/></a> 
<a href="/catalog/12356/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/6163-860-01.jpg"/></a> 
<a href="/catalog/12357/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9240-36-01.jpg"/></a> 
<a href="/catalog/12358/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9240-478-01.jpg"/></a> 
<a href="/catalog/12359/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9240-610-01.jpg"/></a> 
<a href="/catalog/12360/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9420-126-777-778-01.jpg"/></a> 
<a href="/catalog/12361/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9420-127-836-01.jpg"/></a> 
<a href="/catalog/12362/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9420-31-595-01.jpg"/></a> 
<a href="/catalog/12363/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9420-685-673-773-01.jpg"/></a> 
<a href="/catalog/12364/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9420-774-395-01.jpg"/></a> 
<a href="/catalog/12365/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9570-31-01.jpg"/></a> 
<a href="/catalog/12366/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9570-48-01.jpg"/></a> 
<a href="/catalog/12367/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9570-683-01.jpg"/></a> 
<a href="/catalog/12368/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9570-760-01.jpg"/></a> 
<a href="/catalog/12369/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9570-795-01.jpg"/></a> 
<a href="/catalog/12370/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9740-48-01.jpg"/></a> 
<a href="/catalog/12371/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9740-67-01.jpg"/></a> 
<a href="/catalog/12372/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9740-783-01.jpg"/></a> 
<a href="/catalog/12373/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9740-786-01.jpg"/></a> 
<a href="/catalog/12374/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9740-839-01.jpg"/></a> 
<a href="/catalog/12375/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9820-774-01.jpg"/></a> 
<a href="/catalog/12376/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9820-837-01.jpg"/></a> 
<a href="/catalog/12377/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9820-860-01.jpg"/></a> 
<a href="/catalog/12378/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9860-475-01.jpg"/></a> 
<a href="/catalog/12379/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9860-478-01.jpg"/></a> 
<a href="/catalog/12380/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9860-48-01.jpg"/></a> 
<a href="/catalog/12381/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9860-565-01.jpg"/></a> 
<a href="/catalog/12382/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/9860-610-01.jpg"/></a> 
<a href="/catalog/12351/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/102-36-01.jpg"/></a> 
<a href="/catalog/12352/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/102-48-01.jpg"/></a> 
<a href="/catalog/12353/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/102-67-01.jpg"/></a> 
<a href="/catalog/12354/"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_04_25/6163-774-01.jpg"/></a>
</div>',
    'news_images/2023_04_25.jpg'
  );

INSERT INTO
  news (title, data, image_path)
VALUES
  (
    'Новое поступление',
    '<p>Качественные, стильные - размеры 21-25</p>
<p>С 14 июня в продаже (фото кликабельны - можно посмотреть цены и наличие):</p>
<div>
<a href="/catalog/9222"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_06_14/6163-48-01.jpg"/></a> 
<a href="/catalog/12354"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_06_14/6163-774-01.jpg"/></a> 
<a href="/catalog/12383"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_06_14/6163-776-01.jpg"/></a> 
<a href="/catalog/12384"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_06_14/6163-778-01.jpg"/></a> 
<a href="/catalog/12385"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_06_14/6163-781-01.jpg"/></a> 
</div>',
    'news_images/2023_06_14/Fotoram.io.jpg'
  );

INSERT INTO
  news (title, data, image_path)
VALUES
  (
    'Скидки на сандалии',
    'С 1 по 13 июля 2023 года действуют скидки - 20% на покупку сандалий. Если 2 или более пары - 25% ',
    'news_images/2023_07_03.png'
  );

INSERT INTO
  news (title, data, image_path)
VALUES
  (
    'Школьные туфли',
    'Новое поступление школьных туфель для девочек, размеры 30-36',
    'news_images/2023-08-15.jpg'
  );

INSERT INTO
  news (title, data, image_path)
VALUES
  (
    'Зима скоро',
    '
 <p>Теплые, мягкие, клевые - размеры 22-36</p>
<p>С 12 октября в продаже (фото кликабельны - можно посмотреть цены и наличие):</p>
<div>
<a href="/catalog/12409"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_10_11/9195-31 - 01.jpg"/></a> 
<a href="/catalog/12403"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_10_11/1350-31 - 01.jpg"/></a> 
<a href="/catalog/12404"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_10_11/1350-32 - 01.jpg"/></a> 
<a href="/catalog/12405"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_10_11/1350-48 - 01.jpg"/></a> 
<a href="/catalog/12407"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_10_11/1350-593 - 01.jpg"/></a> 
<a href="/catalog/12414"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_10_11/9195-109 - 01.jpg"/></a> 
<a href="/catalog/12406"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_10_11/1350-576 - 01.jpg"/></a>
<a href="/catalog/12408"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_10_11/1350-752 - 01.jpg"/></a>
<a href="/catalog/12410"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_10_11/9195-32 - 01.jpg"/></a>
<a href="/catalog/12411"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_10_11/9195-48 - 01.jpg"/></a>
<a href="/catalog/12412"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_10_11/9195-576 - 01.jpg"/></a>
<a href="/catalog/12413"><img style="margin: 5px" src="https://cipo.kz/static/news_images/2023_10_11/9195-752 - 01.jpg"/></a>
</div>',
    'news_images/news_images/2023_10_11/Fotoram.io.jpg'
  );

INSERT INTO
  news (title, data, image_path)
VALUES
  (
    'Зимние скидки',
    'C 1 января по 1 февраля 2024 года стартовали приятные праздничные скидки: 
Зимняя коллекция от 30% до 50% 
Лето - весна до 20% 
Покупайте с выгодой, поспешите за покупками!
',
    'news_images/2024-01-11.jpg'
  );

INSERT INTO
  news (title, data, image_path)
VALUES
  (
    'Готовь весеннюю-летнюю обувь зимой',
    'Более 30 новых моделей - уже сейчас.
<a href="/catalog/" style="color:blue">Перейти в каталог</a>',
    'news_images/2024_03_15/Fotoram.io (3).jpg'
  );

END -- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT
  'down SQL query';

-- +goose StatementEnd