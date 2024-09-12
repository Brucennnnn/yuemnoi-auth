package route

// func ItemsRouter(app *fiber.App, db *gorm.DB) {
// 	// Get all items
// 	app.Get("/items", func(c *fiber.Ctx) error {
// 		var items []model.Item
// 		if err := db.Find(&items).Error; err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"error": "Unable to retrieve items",
// 			})
// 		}
// 		return c.JSON(items)
// 	})

// 	// Create a new item
// 	app.Post("/items", func(c *fiber.Ctx) error {
// 		var item model.Item
// 		if err := c.BodyParser(&item); err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"error": "Invalid request payload",
// 			})
// 		}

// 		if err := db.Table("items").Create(&item).Error; err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"error": "Unable to create item",
// 			})
// 		}

// 		return c.Status(fiber.StatusCreated).JSON(item)
// 	})

// 	// Update an item by ID
// 	app.Put("/items/:id", func(c *fiber.Ctx) error {
// 		id := c.Params("id")
// 		var item model.Item
// 		if err := db.First(&item, id).Error; err != nil {
// 			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 				"error": "Item not found",
// 			})
// 		}

// 		if err := c.BodyParser(&item); err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"error": "Invalid request payload",
// 			})
// 		}

// 		if err := db.Save(&item).Error; err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"error": "Unable to update item",
// 			})
// 		}

// 		return c.JSON(item)
// 	})

// 	// Delete an item by ID
// 	app.Delete("/items/:id", func(c *fiber.Ctx) error {
// 		id := c.Params("id")
// 		if err := db.Delete(&model.Item{}, id).Error; err != nil {
// 			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 				"error": "Item not found",
// 			})
// 		}

// 		return c.SendStatus(fiber.StatusNoContent)
// 	})
// }
