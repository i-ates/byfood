import React, { useState, useEffect } from "react";
import axios from "axios";
import "./Tech.css";

export function UserList() {
    const [users, setUsers] = useState([]);

    // Fetch users from the API on component mount
    useEffect(() => {
        axios.get(`${process.env.REACT_APP_API_URL}/api/get-users`)
            .then(resp => setUsers(resp.data));
    }, []); // Empty dependency array makes this effect run once after initial render

    return (
        <ul className="users">
            {/* Map over the users and render each user as a list item */}
            {users.map((user, i) => (
                <li key={i}>
                    <b>{user.id}</b>: {user.name}
                </li>
            ))}
        </ul>
    );
}
