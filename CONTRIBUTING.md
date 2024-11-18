---

### `CONTRIBUTING.md`:

```markdown
# Contributing to MoneyWise

Thank you for considering contributing to MoneyWise! We appreciate your efforts in making this project better. Please follow the guidelines below to ensure a smooth development process.

## Branching Strategy

Use **Git Flow** or a **Simplified Workflow** to manage branches and keep things structured. Here's how you can apply it to this project:

### 1. Main Branches
- **`main`** (or **`master`**):
  - This is your production-ready branch. All stable features and code are merged here.
  - Every commit to `main` should be thoroughly tested, with code that works in production.
- **`develop`**:
  - The `develop` branch is where features are integrated and tested before being merged into `main`.
  - It reflects the state of the project in a working (but not yet production-ready) state.

### 2. Feature Branches
Feature branches are created for new functionalities, improvements, or bug fixes.

- **Branch naming**:
  - `feature/{feature-name}`: For new features.
  - `bugfix/{issue-name}`: For bug fixes.
  - `refactor/{refactor-name}`: For code refactoring or optimization.

- **Workflow**:
  1. Create a branch off of `develop`:
     ```bash
     git checkout -b feature/your-feature-name
     ```
  2. Implement the feature, commit frequently, and push to your remote branch.
  3. Once the feature is completed, open a Pull Request (PR) from the feature branch to `develop`.

### 3. Hotfix Branches
If a critical issue arises in production, use hotfix branches:
- **Branch naming**:
  - `hotfix/{issue-name}`: For critical production fixes.
  
- **Workflow**:
  1. Create a branch off `main`:
     ```bash
     git checkout -b hotfix/your-hotfix-name
     ```
  2. Fix the issue, commit changes, and open a PR to merge back to both `main` and `develop`.

## Commit Message Convention

Use clear, consistent commit messages to describe the changes you're making. A good structure to follow is:

- **Type of Change**: Use conventional commit types to indicate the change (feature, fix, chore, refactor, etc.).
- **Brief Summary**: One line summarizing the change.
- **Detailed Description** (optional): A longer explanation of the change, why it was made, and how.

### Commit Message Format:
<type>(<scope>): <summary>

[optional body]


### Examples:
- `feat(auth): add login functionality`
- `fix(transaction): resolve transaction sorting bug`
- `refactor(user): optimize user data fetching`
- `chore(tests): add unit tests for transaction service`

### Types of Changes:
- **feat**: A new feature for the app (e.g., a new microservice or new functionality).
- **fix**: A bug fix.
- **chore**: Routine tasks like setting up configurations, package updates, or testing.
- **refactor**: Code changes that neither fix bugs nor add features but improve the code structure.
- **docs**: Documentation changes.

## Pull Requests
1. **PR Review**: All pull requests should be reviewed before merging.
2. **Testing**: Ensure that all tests pass before creating a pull request.
3. **PR Description**: Provide a clear description of what your PR addresses and reference the relevant issue number (if applicable).
4. **Merging**: Merge your PR only after approval and ensure the branch is up to date with `develop` before merging.

---

Thank you for your contributions!
