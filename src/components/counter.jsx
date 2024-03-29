const { useState } = React;

export default function Counter() {
    const [count, setCount] = useState(0);

    const incrementCount = () => {
        setCount(count + 1);
    };

    return (
        <div className="boxed">
            <h2> Counter component — all client side </h2>
            <p>You clicked the button {count} times.</p>
            <button onClick={incrementCount}>Click me</button>
        </div>
    );
}  