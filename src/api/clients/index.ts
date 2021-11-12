import express from "express";
const router = express.Router();
const clients = [
  {"id":"0", "dni":"1234567t", "name": "Carmelo", "type_diet": "Vegan"},
  {"id":"1", "dni":"7654321e", "name": "Manola", "type_diet": "Buena de boca"}
];

router.get("/", (req, res) => {
  return res.send(clients);
});

module.exports = router;
