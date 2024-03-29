const { useState, useEffect } = React;

export default function Time() {
    const [currentTime, setCurrentTime] = useState('');

    const fetchTime = () => {
        fetch('/api/time')
            .then(response => response.json())
            .then(data => setCurrentTime(data.time))
            .catch(error => console.error("There was an error fetching the time: ", error));
    };

    useEffect(() => {
        fetchTime();
    }, []);

    return (
        <div className="boxed">
            <h2> Time component — API driven </h2>
            <p>Current time: {currentTime}</p>
            <button onClick={fetchTime}>Fetch Time</button>
        </div>
    );
}
