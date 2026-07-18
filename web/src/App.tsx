function App() {
  return (
    <div className="min-h-screen bg-slate-100">
      <nav className="bg-blue-900 text-white px-8 py-4 flex justify-between">
        <h1 className="text-2xl font-bold">VEGO NETWORK</h1>

        <div className="flex gap-6">
          <a href="#">Home</a>
          <a href="#">Platform</a>
          <a href="#">Projects</a>
          <a href="#">Invest</a>
          <a href="#">Contact</a>
        </div>
      </nav>

      <section className="text-center py-24 px-6">
        <h2 className="text-6xl font-bold text-blue-900">
          Building Africa's Digital Agricultural Economy
        </h2>

        <p className="text-xl text-gray-600 mt-8 max-w-3xl mx-auto">
          Connecting Farmers, Investors, Processors, Vendors and
          Agro-Entrepreneurs through technology and tokenized investments.
        </p>

        <div className="mt-12 flex justify-center gap-6">
          <button className="bg-blue-900 text-white px-8 py-4 rounded-lg hover:bg-blue-700">
            Become an Investor
          </button>

          <button className="bg-emerald-600 text-white px-8 py-4 rounded-lg hover:bg-emerald-500">
            Register as Farmer
          </button>
        </div>
      </section>
    </div>
  );
}

export default App;