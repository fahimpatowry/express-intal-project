import app from "./app";
import "dotenv/config";
import connectDB from "./config/db";

const PORT = process.env.PORT || 4000;

const startServer = async () => {
  try {
    await connectDB();

    app.listen(PORT, () => {
      console.log(`🚀 Server running on ${PORT}`);
    });
  } catch (error) {
    console.error("❌ Failed to start server", error);
    process.exit(1);
  }
};

startServer();
