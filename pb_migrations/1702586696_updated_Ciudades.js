/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("n6ikmizd6its0lt")

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "dboivxgw",
    "name": "Nombre",
    "type": "text",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "min": null,
      "max": null,
      "pattern": ""
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("n6ikmizd6its0lt")

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "dboivxgw",
    "name": "Quito",
    "type": "text",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "min": null,
      "max": null,
      "pattern": ""
    }
  }))

  return dao.saveCollection(collection)
})
