document.addEventListener('DOMContentLoaded', () => {
    const urlParams = new URLSearchParams(window.location.search);
    const productName = urlParams.get('name');
    if (productName) {
        fetchProductDetails(productName);
    } else {
        console.error('Nome do produto não fornecido');
    }
});

function fetchProductDetails(productName) {
    const url = `http://localhost:8080/search?keyword=${encodeURIComponent(productName)}`;

    fetch(url)
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .then(products => {
            console.log('Produtos retornados:', products); // Log para depuração
            if (products.length > 0) {
                const product = products[0]; // Supondo que o primeiro produto da lista é o que queremos
                document.title = product.name; // Define o título da página com o nome do produto
                const productTitle = document.getElementById('product-title');
                productTitle.textContent = product.name; // Define o título no header com o nome do produto
                renderProductDetails(product);
            } else {
                console.error('Produto não encontrado');
                renderNotFound();
            }
        })
        .catch(error => console.error('Erro ao buscar detalhes do produto:', error));
}

function renderProductDetails(product) {
    const productDetails = document.getElementById('product-details');
    
    productDetails.innerHTML = `
        <img src="${product.imageURL || 'placeholder.jpg'}" alt="${product.name}">
        <h2>${product.name}</h2>
        <p>${product.description || 'Descrição não disponível'}</p>
        <p><strong>Preço:</strong> $${product.price}</p>
        <p><strong>Categoria:</strong> ${product.category}</p>
        <p><strong>Avaliação:</strong> ${product.rating}</p>
    `;
}

function renderNotFound() {
    const productDetails = document.getElementById('product-details');
    
    productDetails.innerHTML = `
        <p>Produto não encontrado.</p>
    `;
}
