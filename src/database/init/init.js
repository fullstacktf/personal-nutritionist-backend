db = connect("localhost:27017/nutriguide");

db.createCollection('users');
db.createCollection('recipes');
db.createCollection('events');

db.users.insertMany([
  {
    name: "Axel Tajada Herrera",
    email: "axeltaj@gmail.com",
    username: "Axel Tajada",
    password: "alex",
    role: "Nutricionista",
    dni: "12332286G",
    birthday: "05/05/1989",
    phone: 635711533,
    description: "Especialista en nutrición deportiva, ¿necesitas consejo?,¡llamame!",
    isVerified: true,
    education: ["Graduado en ciencias de la salud y el deporte en la ULL", "Doctorado en asesoramiento dietetico"],
    specialities: ["deportiva"],
    price: 89.99,
    likes: 60
  },
  {
    name: "Judith Reig",
    email: "Judith.Reig@hotmail.com",
    username: "Judith",
    password: "judith",
    role: "Nutricionista",
    dni: "121234386P",
    birthday: "15/08/1997",
    phone: 637895334,
    description: "Especialista en nutrición vegetariana",
    isVerified: true,
    education: ["Graduado en ciencias de la salud"],
    specialities: ["vegetariana"],
    price: 75,
    likes: 120
  },
  {
    name: "Aitana Reyes Negrin",
    email: "aitanarn@hotmail.com",
    username: "Aitana Reyes",
    password: "aitana",
    role: "Nutricionista",
    dni: "121287649Q",
    birthday: "01/01/2000",
    phone: 680322464,
    description: "Asesoramiento nutricional enfocado a trastornos alimenticios",
    isVerified: false,
    specialities: ["Trastornos alimenticios"],
    price: 55,
    likes: 47
  },
  {
    name: "Casilda Corral Lucena",
    email: "casicorrlu@hotmail.com",
    username: "Casilda Corral",
    password: "casilda",
    role: "Cliente",
    dni: "103948594A",
    birthday: "22/02/2007",
    phone: 679897074,
    description: "",
    weight: 65,
    height: 155,
    typeDiet: "Vegetariana",
    intolerances: ["nueces", "Lactosa"],
  },
  {
    name: "Martín Paredes Cantos",
    email: "martinparedes@hotmail.com",
    username: "Martin Paredes",
    password: "martin",
    role: "Cliente",
    dni: "10399984B",
    birthday: "16/09/2005",
    phone: 679897674,
    description: "Me gusta comer 🍕",
    weight: 140,
    height: 165,
    typeDiet: "Omnívora",
  },
]);

db.events.insertMany([
  {
    title: "Reunión con Martin Paredes",
    status: "Aceptado",
    description: "Asesoramiento sobre malos hábitos",
    startingDate: "25/04/2022 13:00",
    endingDate: "25/04/2022 13:45"
  },
  {
    title: "Reunión con Casilda",
    status: "Rechazado",
    description: "Revisión mensual",
    startingDate: "22/03/2022 12:00",
    endingDate: "22/03/2022 12:30"
  },
  {
    title: "Reunión con Martin Paredes",
    status: "Pendiente",
    description: "Revisión mensual",
    startingDate: "22/05/2022 09:00",
    endingDate: "22/05/2022 09:30"
  }
]);

db.recipes.insertMany([
  {
    name: "Papas con carne 🥔🍖",
    typeMeal: "Almuerzo",
    date: "03/02/2021",
    typeDiet: "Omnívora",
    alergens: ["Dioxido de azufre y sulfitos"],
    ingredients: ["Papas", "Carne", "Sal", "Vino"],
    preparation: "1º Pelar las papas. 2º Hervir agua. 3º Cocinar las papas y la carne. 4º Agregar un chorro de vino a la carne",
    coockingTime: "60 min"
  },
  {
    name: "Crema de puerros 🥬🍲",
    typeMeal: "Cena",
    date: "03/02/2021",
    typeDiet: "Vegana",
    alergens: ["Dioxido de azufre y sulfitos"],
    ingredients: ["4 Puerros", "2 papas", "250ml de agua", "200ml de leche", "1 cebolla", "50ml aceite de oliva virgen extra", "sal"],
    preparation: "1º Cortar puerros. 2º Hervir agua. 3º Cocinar las papas. 4º Agregar aceite y sal",
    coockingTime: "40 min"
  },
  {
    name: "Ostras al limón 🦪🍋",
    typeMeal: "Cena",
    date: "03/02/2021",
    typeDiet: "Omnívora",
    alergens: ["Moluscos"],
    ingredients: ["4 moluscos", "1 Limon"],
    preparation: "1º Hervir agua. 2º Cocinar las otras. 3º Agregar limón",
    coockingTime: "45 min"
  },
]);
