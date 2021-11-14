import express from "express";

const router = express.Router();
const users = [
  {
    "id": 1,
    "email": "lalolita@gmail.com",
    "username": "Lola",
    "role": "client",
    "password": "lalola1234"
  },
  {
    "id": 2,
    "email": "elpanadero@gmail.com",
    "username": "Juan",
    "role": "client",
    "password": "elpanesmipasion"
  }
];

router.get("/", (req, res) => {
  res.json(users);
});

router.get("/:id", (req, res) => {
  const id = req.params.id;
  const user = users.find(u => u.id == parseInt(id));
  return res.send(user);
});

router.post("/", (req, res) => {
  const newUser = req.body;
  users.push(newUser);
  res.json(newUser);
});

router.put("/:id", (req, res) => {
  const id = req.params.id;
  const newUser = req.body;
  const userIndex = users.findIndex(u => u.id == parseInt(id));
  users[userIndex] = newUser;
  res.json(users[userIndex]);
});

router.delete("/:id", (req, res) => {
  const id = req.params.id;
  const userIndex = users.findIndex(u => u.id == parseInt(id));
  const deletedUser = users.splice(userIndex, 1);
  return res.send(deletedUser);
});

module.exports = router;
