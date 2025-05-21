import { Carousel } from "antd";
function Carousell() {
  const contentStyle = {
    height: "220px",
    color: "#fff",
    lineHeight: "220px",
    textAlign: "center",
    background: "#364d79",
  };

  return (
    <Carousel autoplay>
      <div>
        <h3 style={contentStyle}>1</h3>
      </div>
      <div>
        <h3 style={contentStyle}>2</h3>
      </div>
      <div>
        <h3 style={contentStyle}>3</h3>
      </div>
      <div>
        <h3 style={contentStyle}>4</h3>
      </div>
    </Carousel>
  );
}

export default Carousell;
