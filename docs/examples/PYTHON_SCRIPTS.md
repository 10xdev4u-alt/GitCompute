# 🐍 Example: Running Python Scripts

You can use GitCompute to run Python data processing scripts on the cloud.

## Scenario
You have a script `process_data.py` that takes 10 minutes to run. You don't want to heat up your laptop.

## Prerequisite
Your `worker.yml` must checkout the code. Ensure the `actions/checkout` step is present.

## Execution

```bash
git-compute run \
  --cmd "python3 scripts/process_data.py" \
  --os ubuntu-latest
```

## Dependencies
If you need libraries like `pandas` or `numpy`:

```bash
git-compute run \
  --cmd "pip install pandas numpy && python3 scripts/process_data.py" \
  --os ubuntu-latest
```

## Matrix Strategy (Manual)
To process 3 different datasets in parallel, open 3 terminal tabs:

**Tab 1:**
```bash
git-compute run --cmd "python3 process.py --dataset A"
```

**Tab 2:**
```bash
git-compute run --cmd "python3 process.py --dataset B"
```

**Tab 3:**
```bash
git-compute run --cmd "python3 process.py --dataset C"
```

Result: **3x Speedup**.

```