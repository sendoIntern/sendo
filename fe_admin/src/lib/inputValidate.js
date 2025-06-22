export const inputValidate = {
  isValidName: (name) => {
    const nameRegex = /^[a-zA-Z\s]{2,20}$/;
    return nameRegex.test(name);
  },
  isValidQuantity: (quantity) => {
    const quantityRegex = /^[1-9]\d*$/; // Positive integers only
    return quantityRegex.test(quantity);
  },
  isValidPrice: (price) => {
    const priceRegex = /^\d+(\.\d{1,2})?$/; // Non-negative numbers with up to two decimal places
    return priceRegex.test(price);
  },
};
