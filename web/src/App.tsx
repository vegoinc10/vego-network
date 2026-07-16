function App() {
  return (
    <div
      style={{
        minHeight: "100vh",
        backgroundColor: "#F8FAFC",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        textAlign: "center",
        padding: "20px",
      }}
    >
      <div>
        <h1 style={{ color: "#1E3A8A", fontSize: "48px" }}>
          VEGO NETWORK
        </h1>

        <h2 style={{ color: "#10B981" }}>
          Building Africa's Digital Agricultural Economy
        </h2>

        <p style={{ maxWidth: "700px", margin: "20px auto" }}>
          Connecting Farmers, Investors, Processors, Vendors and Agro-Entrepreneurs
          through technology, finance and innovation.
        </p>

        <button
          style={{
            backgroundColor: "#1E3A8A",
            color: "white",
            padding: "15px 30px",
            border: "none",
            borderRadius: "8px",
            cursor: "pointer",
            fontSize: "18px",
          }}
        >
          Get Started
        </button>
      </div>
    </div>
  );
}

export default App;