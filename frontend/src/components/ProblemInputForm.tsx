import { useState } from 'react';
import ApiKeyManager from './ApiKeyManager';

interface ProblemInputFormProps {
  onSubmit: (data: {
    problemText: string;
    language: string;
    userLevel: string;
    provider: string;
    apiKey: string;
  }) => void;
  isLoading: boolean;
}

export default function ProblemInputForm({ onSubmit, isLoading }: ProblemInputFormProps) {
  const [problemText, setProblemText] = useState<string>('');
  const [language, setLanguage] = useState<string>('javascript');
  const [userLevel, setUserLevel] = useState<string>('intermediate');
  const [provider, setProvider] = useState<string>('openai');
  const [apiKey, setApiKey] = useState<string>('');
  const [showApiKeyManager, setShowApiKeyManager] = useState<boolean>(false);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSubmit({
      problemText,
      language,
      userLevel,
      provider,
      apiKey,
    });
  };

  const handleApiKeyChange = (newApiKey: string) => {
    setApiKey(newApiKey);
  };

  return (
    <div className="bg-secondary text-white p-6 rounded-lg shadow-md">
      <form onSubmit={handleSubmit}>
        <div className="mb-4">
          <label htmlFor="problem" className="block text-sm font-medium text-primary mb-1">
            LeetCode Problem
          </label>
          <textarea
            id="problem"
            value={problemText}
            onChange={(e) => setProblemText(e.target.value)}
            placeholder="Paste your LeetCode problem here..."
            className="w-full h-40 p-3 border border-gray-700 rounded-md bg-secondary-light text-white"
            required
          />
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
          <div>
            <label htmlFor="language" className="block text-sm font-medium text-primary mb-1">
              Programming Language
            </label>
            <select
              id="language"
              value={language}
              onChange={(e) => setLanguage(e.target.value)}
              className="w-full p-2 border border-gray-700 rounded-md bg-secondary-light text-white"
            >
              <option value="javascript">JavaScript</option>
              <option value="python">Python</option>
              <option value="java">Java</option>
              <option value="cpp">C++</option>
              <option value="go">Go</option>
              <option value="typescript">TypeScript</option>
              <option value="ruby">Ruby</option>
              <option value="swift">Swift</option>
              <option value="kotlin">Kotlin</option>
              <option value="rust">Rust</option>
            </select>
          </div>

          <div>
            <label htmlFor="level" className="block text-sm font-medium text-primary mb-1">
              Your Experience Level
            </label>
            <select
              id="level"
              value={userLevel}
              onChange={(e) => setUserLevel(e.target.value)}
              className="w-full p-2 border border-gray-700 rounded-md bg-secondary-light text-white"
            >
              <option value="beginner">Beginner</option>
              <option value="intermediate">Intermediate</option>
              <option value="advanced">Advanced</option>
            </select>
          </div>
        </div>

        <div className="mb-4">
          <label htmlFor="provider" className="block text-sm font-medium text-primary mb-1">
            AI Provider
          </label>
          <select
            id="provider"
            value={provider}
            onChange={(e) => {
              setProvider(e.target.value);
              setShowApiKeyManager(true);
            }}
            className="w-full p-2 border border-gray-700 rounded-md bg-secondary-light text-white"
          >
            <option value="openai">OpenAI (GPT-4)</option>
            <option value="gemini">Google Gemini</option>
            <option value="claude">Anthropic Claude</option>
            <option value="groq">Groq (Llama 3)</option>
          </select>
        </div>

        {showApiKeyManager ? (
          <ApiKeyManager provider={provider} onApiKeyChange={handleApiKeyChange} />
        ) : (
          <div className="mb-6">
            <div className="flex items-center">
              <input
                id="saveApiKey"
                type="checkbox"
                className="h-4 w-4 text-primary border-gray-700 rounded"
                checked={true}
                onChange={() => setShowApiKeyManager(true)}
              />
              <label htmlFor="saveApiKey" className="ml-2 block text-sm text-white">
                Save API key locally (never sent to our servers)
              </label>
            </div>
            <input
              type="password"
              value={apiKey}
              onChange={(e) => setApiKey(e.target.value)}
              placeholder={`Enter your ${provider} API key`}
              className="w-full mt-2 p-2 border border-gray-700 rounded-md bg-secondary-light text-white"
              required
            />
          </div>
        )}

        <button
          type="submit"
          disabled={isLoading}
          className={`w-full py-3 px-4 rounded-md text-secondary font-bold ${
            isLoading
              ? 'bg-primary-dark cursor-not-allowed'
              : 'bg-primary hover:bg-primary-light'
          }`}
        >
          {isLoading ? 'Generating Solution...' : 'Solve Problem'}
        </button>
      </form>
    </div>
  );
}
