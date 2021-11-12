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

router.post("/", (req, res) => {
  const newNutritionist = req.body;
  nutritionists.push(newNutritionist);
  res.json(newNutritionist);
});

router.put("/:id", (req, res) => {
  const id = req.params.id;
  const newNutritionist = req.body;
  const nutritionistIndex = nutritionists.findIndex(n => n.id == parseInt(id));
  nutritionists[nutritionistIndex].name = newNutritionist.name;
  res.json(nutritionists[nutritionistIndex]);
});

router.delete("/:id", (req, res) => {
  const id = req.params.id;
  const nutritionistIndex = nutritionists.findIndex(n => n.id == parseInt(id));
  const deletedNutritionist = nutritionists.splice(nutritionistIndex, 1);
  return res.send(deletedNutritionist);
});

module.exports = router;
