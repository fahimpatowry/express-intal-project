import mongoose from "mongoose";

const connectDB = async (): Promise<void> => {
  try {
    const mongoUri =
      process.env.MONGO_URI || "mongodb://localhost:27017/initial-projects";

      console.log(mongoUri)

    await mongoose.connect(mongoUri);

    console.log("✅ MongoDB connected");
  } catch (error) {
    console.error("❌ MongoDB connection failed", error);
    process.exit(1);
  }
};

export default connectDB;
