// API client for interacting with the backend
export const apiClient = {
  // Base URL for the API
  baseUrl: 'http://localhost:8080',

  // Get list of available providers
  async getProviders() {
    try {
      const response = await fetch(`${this.baseUrl}/api/providers`);
      if (!response.ok) {
        throw new Error('Failed to fetch providers');
      }
      return await response.json();
    } catch (error) {
      console.error('Error fetching providers:', error);
      throw error;
    }
  },

  // Solve a LeetCode problem
  async solveProblem(data) {
    try {
      const response = await fetch(`${this.baseUrl}/api/solve`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          problem_text: data.problemText,
          language: data.language,
          user_level: data.userLevel,
          provider: data.provider,
          api_key: data.apiKey,
        }),
      });

      const result = await response.json();
      if (!response.ok) {
        throw new Error(result.error || 'Failed to generate solution');
      }

      return result;
    } catch (error) {
      console.error('Error solving problem:', error);
      throw error;
    }
  },
};

// API key storage utilities
export const apiKeyStorage = {
  // Save API key to local storage
  saveApiKey(provider, apiKey) {
    if (provider && apiKey) {
      localStorage.setItem(`leetcode-helper-${provider}-api-key`, apiKey);
      return true;
    }
    return false;
  },

  // Get API key from local storage
  getApiKey(provider) {
    if (!provider) return null;
    return localStorage.getItem(`leetcode-helper-${provider}-api-key`);
  },

  // Remove API key from local storage
  removeApiKey(provider) {
    if (!provider) return false;
    localStorage.removeItem(`leetcode-helper-${provider}-api-key`);
    return true;
  },

  // Check if API key exists for provider
  hasApiKey(provider) {
    return !!this.getApiKey(provider);
  },
};
