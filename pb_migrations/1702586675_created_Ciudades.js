/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const collection = new Collection({
    "id": "n6ikmizd6its0lt",
    "created": "2023-12-14 20:44:35.026Z",
    "updated": "2023-12-14 20:44:35.026Z",
    "name": "Ciudades",
    "type": "base",
    "system": false,
    "schema": [
      {
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
      }
    ],
    "indexes": [],
    "listRule": null,
    "viewRule": null,
    "createRule": null,
    "updateRule": null,
    "deleteRule": null,
    "options": {}
  });

  return Dao(db).saveCollection(collection);
}, (db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("n6ikmizd6its0lt");

  return dao.deleteCollection(collection);
})
