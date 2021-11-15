import express from "express";
const router = express.Router();
const clients = [
  {"id": 0, "dni": "1234567t", "name": "Carmelo", "type_diet": "Vegan"},
  {"id": 1, "dni": "7654321e", "name": "Manola", "type_diet": "Buena de boca"}
];

router.get("/", (req, res) => {
  return res.send(clients);
});

router.get("/:id", (req, res) => {
  const client = clients.find(c => c.id == parseInt(req.params.id));
  return res.send(client);
});

router.post("/", (req, res) => {
  clients.push(req.body);
  res.json(clients);
});

router.put("/:id", (req, res) => {
  const index = clients.findIndex(c => c.id == parseInt(req.params.id));
  clients[index] = req.body;
  res.json(clients[index]);
});

router.delete("/:id", (req, res) => {
  const index = clients.findIndex(c => c.id == parseInt(req.params.id));
  res.json(clients.splice(index, 1));
});

module.exports = router;
