<template>
  <div>
    <h1>JSON Data</h1>
    <div v-if="jsonData.length">
      <div v-for="(item, index) in jsonData" :key="index" class="item-box">
        <p>
          <strong> ID : {{ index + 1 }} </strong>
        </p>
        <p><strong>Name:</strong> {{ item.name }}</p>
        <p><strong>Description:</strong> {{ item.descri }}</p>
        <p><strong>Boolean:</strong> {{ item.bool }}</p>
      </div>
    </div>
    <button @click="fetchJsonData">Fetch JSON Data</button>
    <button @click="goBack">Go Back</button>

    <!-- 削除するところ -->
    <div class="form-container">
      <h3>Enter Name to Delete</h3>
      <form @submit.prevent="deleteForm">
        <label for="digit">Name</label>
        <input
          type="number"
          id="digit"
          v-model="digit"
          :min="1"
          :max="maxId"
          required
          class="digit-input"
        />
        <button type="submit" class="delete-button">delete</button>
      </form>
      <p v-if="deleteDigit !== null">You entered: {{ deleteDigit }}</p>
    </div>
    <!-- 追加するところ -->
    <h1>Submit Data</h1>
    <form @submit.prevent="submitData">
      <label for="name">Name:</label>
      <input type="number" id="name" v-model="formData.Name" required />

      <label for="descri">Description:</label>
      <input type="text" id="descri" v-model="formData.Descri" required />

      <label for="bool">Boolean:</label>
      <input type="checkbox" id="bool" v-model="formData.Bool" />

      <button type="submit">Submit Data</button>
    </form>
    <p v-if="formError" style="color: red">{{ formError }}</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      formData: {
        Name: null,
        Descri: "",
        Bool: false,
      },
      jsonData: [],
      formError: null, // エラーメッセージ用
      deleteDigit: null,
      digit: null,
    };
  },
  computed: {
    maxId() {
      // 最大IDを計算して返す
      return this.jsonData.length || 0;
    },
  },
  methods: {
    async fetchJsonData() {
      try {
        const response = await fetch("http://localhost:8081/json");
        if (!response.ok) {
          console.log("Response status:", response.status);
          throw new Error("Failed to fetch JSON data");
        }
        const data = await response.json();
        console.log("Fetched data:", data); // デバッグ用
        if (Array.isArray(data)) {
          this.jsonData = data;
        } else {
          console.error("Data is not an array");
        }
      } catch (error) {
        console.error("Error fetching JSON data:", error);
      }
    },
    async submitData() {
      // フォームのバリデーション
      if (!this.formData.Name || !this.formData.Descri) {
        this.formError = "All fields are required.";
        return;
      }

      try {
        const response = await fetch(`http://localhost:8081/add`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(this.formData),
        });
        if (!response.ok) {
          throw new Error("Failed to submit data");
        }
        alert("Data submitted successfully!");
        this.formError = null; // エラーメッセージをクリア
        this.formData = { Name: null, Descri: "", Bool: false }; // フォームをリセット
      } catch (error) {
        console.error("Error submitting data:", error);
      }
    },

    async deleteForm() {
      // フォームのバリデーション
      if (!this.digit) {
        this.deleteDigit = null;
        return;
      }
      console.log("Deleting data with ID:", this.digit); // デバッグ用
      try {
        const response = await fetch(
          `http://localhost:8081/delete/${this.digit}`,
          {
            method: "DELETE",
          }
        );
        console.log("Response status:", response.status, this.digit); // デバッグ用
        if (!response.ok) {
          throw new Error("Failed to delete data");
        }
        alert("Data deleted successfully!");
        this.deleteDigit = this.digit; // 削除したIDを表示
        this.digit = null; // フォームをリセット
      } catch (error) {
        console.error("Error deleting data:", error);
      }
    },
    goBack() {
      this.$router.push("/");
    },
  },
};
</script>

<style scoped>
.item-box {
  border: 1px solid #ccc; /* ボーダーの色 */
  padding: 10px; /* 内側の余白 */
  margin: 10px 0; /* 外側の余白 */
  border-radius: 5px; /* 角の丸み（オプション） */
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1); /* 影（オプション） */
}

.form-container {
  width: 300px; /* フォームの幅を設定 */
  margin: auto;
  padding: 20px;
  border: 0.5px solid #ccc;
  border-radius: 5px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.digit-input {
  width: 10%; /* フィールドの幅を調整 */
  padding: 10px;
  margin-bottom: 10px;
  margin-top: 5px;
  margin-left: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
  text-align: center;
}

.delete-button {
  width: 100%; /* ボタンの幅を調整 */
  padding: 10px;
  border: none;
  border-radius: 5px;
  background-color: #4caf50;
  color: white;
  font-size: 16px;
  cursor: pointer;
}

.submit-button:hover {
  background-color: #45a049;
}
</style>
