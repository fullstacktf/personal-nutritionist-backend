import express from "express";

const router = express.Router();
const nutritionists = [
  { "id": 1, "name": "Nutricionista 1" },
  { "id": 2, "name": "Nutricionista 2" }
];

router.get("/", (req, res) => {
  res.json(nutritionists);
});

router.get("/:id", (req, res) => {
  const id = req.params.id;
  const nutritionist = nutritionists.find(n => n.id == parseInt(id));
  return res.send(nutritionist);
});

module.exports = router;
