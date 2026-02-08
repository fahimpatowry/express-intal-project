import swaggerJsdoc from "swagger-jsdoc";

const PORT = process.env.PORT || 4000;

const options = {
  definition: {
    openapi: "3.0.0",
    info: {
      title: "E-commerce client",
      version: "1.0.0",
      description: "API documentation",
    },
    servers: [
      {
        url: `http://localhost:${PORT}`, // Update your server URL
      },
    ],
  },
  apis: ["./src/routes/*.ts", "./src/controllers/*.ts"], // Path to your route files
};

const swaggerSpec = swaggerJsdoc(options);

export default swaggerSpec;
