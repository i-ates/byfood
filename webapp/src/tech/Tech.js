import React, { useState } from "react";
import { UserListView } from "./UserListView";
import Button from "react-bootstrap/Button";
import { NewUserScreen } from "./NewUserScreen";
import "./Tech.css";

export function Tech() {
    const [users, setUsers] = useState([]); // State for storing user data

    const [showModal, setShowModal] = useState(false); // State for controlling modal visibility

    const toggleModal = () => {
        setShowModal(!showModal); // Function to toggle modal visibility
    };

    return (
        <div className="tech">
            <h2 className="title">User Management for ByFood</h2>

            {/* Button to open the modal */}
            <td>
                <Button variant="primary" onClick={() => toggleModal()} className="custom-new-button">
                    New User
                </Button>
                {/* Modal for adding a new user */}
                <NewUserScreen show={showModal} onHide={toggleModal} />
            </td>

            {/* Component for displaying the list of users */}
            <UserListView />
        </div>
    );
}
