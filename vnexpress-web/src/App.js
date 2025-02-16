import { useEffect, useState } from "react";

function App() {
  const [articles, setArticles] = useState([]);

  useEffect(() => {
    fetch("http://localhost:6869/articles")
      .then((res) => {
        if (!res.ok) {
          throw new Error(`HTTP error! Status: ${res.status}`);
        }
        return res.json();
      })
      .then((data) => {
        console.log("data:", data);
        setArticles(data);
      })
      .catch((err) => console.error("error when call API:", err));
  }, []);

  return (
    <div style={styles.container}>
      <h1 style={styles.header}>VNExpress</h1>
      <ul style={styles.list}>
        {articles.length > 0 ? (
          articles.map((article, index) => (
            <li key={index} style={styles.card}>
              {article.imageURL && (
                <img
                  src={article.imageURL}
                  alt={article.title}
                  style={styles.image}
                />
              )}
              <h2>
                <a href={article.url} target="_blank" rel="noopener noreferrer" style={styles.link}>
                  {article.title}
                </a>
              </h2>
              <p><strong>Date posted:</strong> {article.publishedDate}</p>
              <p><strong>Category:</strong> {article.category}</p>
              <p><strong>Topic:</strong> {article.subCategory}</p>
              <p>{article.description}</p>
            </li>
          ))
        ) : (
          <p>There are no articles.</p>
        )}
      </ul>
    </div>
  );
}

const styles = {
  container: {
    maxWidth: "800px",
    margin: "20px auto",
    padding: "20px",
    fontFamily: "Arial, sans-serif",
  },
  header: {
    textAlign: "center",
    color: "#333",
  },
  list: {
    listStyle: "none",
    padding: 0,
  },
  card: {
    border: "1px solid #ddd",
    padding: "15px",
    marginBottom: "20px",
    borderRadius: "8px",
    boxShadow: "0 2px 5px rgba(0,0,0,0.1)",
  },
  image: {
    width: "100%",
    maxHeight: "300px",
    objectFit: "cover",
    borderRadius: "8px",
    marginBottom: "10px",
  },
  link: {
    textDecoration: "none",
    color: "#007bff",
    fontSize: "18px",
    fontWeight: "bold",
  },
};

export default App;
