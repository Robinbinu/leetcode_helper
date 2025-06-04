import { useState, useEffect } from 'react';

interface ApiKeyManagerProps {
  provider: string;
  onApiKeyChange: (apiKey: string) => void;
}

export default function ApiKeyManager({ provider, onApiKeyChange }: ApiKeyManagerProps) {
  const [apiKey, setApiKey] = useState<string>('');
  const [isSaved, setIsSaved] = useState<boolean>(false);

  // Load API key from local storage when provider changes
  useEffect(() => {
    const savedKey = localStorage.getItem(`leetcode-helper-${provider}-api-key`);
    if (savedKey) {
      setApiKey(savedKey);
      setIsSaved(true);
      onApiKeyChange(savedKey);
    } else {
      setApiKey('');
      setIsSaved(false);
      onApiKeyChange('');
    }
  }, [provider, onApiKeyChange]);

  const handleSaveApiKey = () => {
    if (apiKey.trim()) {
      localStorage.setItem(`leetcode-helper-${provider}-api-key`, apiKey);
      setIsSaved(true);
      onApiKeyChange(apiKey);
    }
  };

  const handleClearApiKey = () => {
    localStorage.removeItem(`leetcode-helper-${provider}-api-key`);
    setApiKey('');
    setIsSaved(false);
    onApiKeyChange('');
  };

  return (
    <div className="mt-4 p-4 bg-secondary-light border border-secondary rounded-md">
      <h3 className="text-lg font-medium mb-2 text-white">API Key for {provider}</h3>
      <div className="flex gap-2">
        <input
          type="password"
          value={apiKey}
          onChange={(e) => setApiKey(e.target.value)}
          placeholder={`Enter your ${provider} API key`}
          className="flex-1 p-2 border border-gray-300 rounded-md bg-background text-secondary"
        />
        <button
          onClick={handleSaveApiKey}
          className="bg-primary text-secondary px-4 py-2 rounded-md hover:bg-primary-dark font-medium"
        >
          Save
        </button>
        {isSaved && (
          <button
            onClick={handleClearApiKey}
            className="bg-red-600 text-white px-4 py-2 rounded-md hover:bg-red-700"
          >
            Clear
          </button>
        )}
      </div>
      {isSaved && (
        <p className="text-primary text-sm mt-2">
          ✓ API key saved locally (not sent to our servers)
        </p>
      )}
    </div>
  );
}
