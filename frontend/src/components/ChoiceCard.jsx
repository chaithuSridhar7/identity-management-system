function ChoiceCard({ children, onClick, disabled = false }) {
  return (
    <button
      type="button"
      className={`choice-card ${disabled ? "choice-card--disabled" : ""}`}
      onClick={onClick}
      disabled={disabled}
    >
      {children}
    </button>
  );
}

export default ChoiceCard;