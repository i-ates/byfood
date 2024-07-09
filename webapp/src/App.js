import React from 'react';
import './App.css';
import { Tech } from "./tech/Tech";

// The main component of the application
export function App() {
    return (
        <div className="app">
            {/* Video background element */}
            <div className="video-background">
                {/* Video element with autoplay, loop, and muted properties */}
                <video autoPlay loop muted>
                    {/* Video source */}
                    <source src="https://vz-6860d929-9e5.b-cdn.net/a94b3f76-d560-47c1-80b8-7dc92dcca5cc/play_720p.mp4" type="video/mp4" />
                    Your browser does not support the video tag.
                </video>
            </div>
            {/* Render the Tech component */}
            <div>
                <Tech />
            </div>
        </div>
    );
}
