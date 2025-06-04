import { useState } from 'react';
import { apiClient } from '../utils/api';
import ProblemInputForm from './ProblemInputForm';
import SolutionDisplay from './SolutionDisplay';

export default function App() {
  const [solution, setSolution] = useState(null);
  const [error, setError] = useState(null);
  const [isLoading, setIsLoading] = useState(false);

  const handleSubmit = async (data) => {
    setIsLoading(true);
    setError(null);
    setSolution(null);

    try {
      const result = await apiClient.solveProblem(data);
      setSolution(result);
    } catch (err) {
      setError(err.message);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div>
      <ProblemInputForm onSubmit={handleSubmit} isLoading={isLoading} />
      <SolutionDisplay solution={solution} error={error} />
    </div>
  );
}