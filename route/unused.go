package route

// func itemRouter(app *fiber.App, db *gorm.DB) {
// 	// Get all item
// 	app.Get("/item", func(c *fiber.Ctx) error {
// 		var item []model.Item
// 		if err := db.Find(&item).Error; err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"error": "Unable to retrieve item",
// 			})
// 		}
// 		return c.JSON(item)
// 	})

// 	// Create a new item
// 	app.Post("/item", func(c *fiber.Ctx) error {
// 		var item model.Item
// 		if err := c.BodyParser(&item); err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"error": "Invalid request payload",
// 			})
// 		}

// 		if err := db.Table("item").Create(&item).Error; err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"error": "Unable to create item",
// 			})
// 		}

// 		return c.Status(fiber.StatusCreated).JSON(item)
// 	})

// 	// Update an item by ID
// 	app.Put("/item/:id", func(c *fiber.Ctx) error {
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
// 	app.Delete("/item/:id", func(c *fiber.Ctx) error {
// 		id := c.Params("id")
// 		if err := db.Delete(&model.Item{}, id).Error; err != nil {
// 			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 				"error": "Item not found",
// 			})
// 		}

// 		return c.SendStatus(fiber.StatusNoContent)
// 	})
// }
