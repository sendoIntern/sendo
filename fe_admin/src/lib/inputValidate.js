export const inputValidate = {
  isValidName: (name) => {
    const nameRegex = /^[\p{L}\s]{2,50}$/u;
    return nameRegex.test(name.trim());
  },

  isValidQuantity: (quantity) => {
    const quantityRegex = /^[1-9]\d*$/; // Positive integers only
    return quantityRegex.test(quantity);
  },

  isValidPrice: (price) => {
    const priceRegex = /^\d+(\.\d{1,2})?$/; // Non-negative numbers with up to two decimal places
    return priceRegex.test(price);
  },

  formatNumberWithCommas: (value) => {
    if (value == null || value === "") return "";
    return value.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ".");
  },

  //cho nhận ký tự 0-9
  parseNumberFromString: (value) => {
    if (!value) return "";
    return value.replace(/[^\d]/g, "");
  },

  //cho nhận các phím Backspace, Delete, ArrowLeft, ArrowRight, Tab
  restrictNumberInputKeys: (e) => {
    const allowedKeys = [
      "Backspace",
      "Delete",
      "ArrowLeft",
      "ArrowRight",
      "Tab",
    ];
    if (!/[0-9]/.test(e.key) && !allowedKeys.includes(e.key)) {
      e.preventDefault();
    }
  },
};
