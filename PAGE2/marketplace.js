document.addEventListener('DOMContentLoaded', () => {
    fetchProducts();
});

function fetchProducts() {
    fetch('http://localhost:8080/search')
        .then(response => response.json())
        .then(data => renderProducts(data))
        .catch(error => console.error('Erro ao buscar produtos:', error));
}

function renderProducts(products) {
    const productList = document.getElementById('product-list');
    productList.innerHTML = '';

    products.forEach(product => {
        const productItem = document.createElement('div');
        productItem.className = 'product-item';

        productItem.innerHTML = `
            <img src="${product.imageUrl}" alt="${product.name}">
            <h2>${product.name}</h2>
            <p>${product.description}</p>
            <p><strong>Price:</strong> $${product.price}</p>
        `;

        productList.appendChild(productItem);
    });
}

