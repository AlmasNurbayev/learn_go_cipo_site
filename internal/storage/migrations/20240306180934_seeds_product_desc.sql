-- +goose Up
-- +goose StatementBegin
INSERT INTO product_desc (id_1c, name_1c, field) VALUES 
('7dac490d-b0d2-11ed-b0f1-50ebf624c538', 'Материал подошвы', 'material_podoshva'), 
('7dac490f-b0d2-11ed-b0f1-50ebf624c538', 'Материал внутри', 'material_inside'), 
('7dac490e-b0d2-11ed-b0f1-50ebf624c538', 'Материал вверх', 'material_up'), 
('28a101cd-b439-11ed-b0f5-50ebf624c538', 'Мальчик/девочка', 'sex'), 
('28a101ce-b439-11ed-b0f5-50ebf624c538', 'Основной цвет', 'main_color'), 
('a001d8e0-a3b3-11ed-b0d2-50ebf624c538', 'ВыгружатьВеб', 'public_web'),
('6c2db24d-a792-11ed-b0d5-50ebf624c538', 'ТоварнаяГруппа', 'product_group'),
('7dac4910-b0d2-11ed-b0f1-50ebf624c538', 'Виды', 'vids')
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE product_desc;
-- +goose StatementEnd
