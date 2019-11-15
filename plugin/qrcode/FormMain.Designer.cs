namespace qrcode
{
    partial class FormMain
    {
        /// <summary>
        /// 必需的设计器变量。
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        /// 清理所有正在使用的资源。
        /// </summary>
        /// <param name="disposing">如果应释放托管资源，为 true；否则为 false。</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows 窗体设计器生成的代码

        /// <summary>
        /// 设计器支持所需的方法 - 不要修改
        /// 使用代码编辑器修改此方法的内容。
        /// </summary>
        private void InitializeComponent()
        {
            System.ComponentModel.ComponentResourceManager resources = new System.ComponentModel.ComponentResourceManager(typeof(FormMain));
            this.pictureBox_qrcode = new System.Windows.Forms.PictureBox();
            this.label_content = new System.Windows.Forms.Label();
            this.textBox_content = new System.Windows.Forms.TextBox();
            this.button_generate = new System.Windows.Forms.Button();
            this.button_clear = new System.Windows.Forms.Button();
            this.button_select = new System.Windows.Forms.Button();
            this.textBox_images = new System.Windows.Forms.TextBox();
            this.label_path = new System.Windows.Forms.Label();
            ((System.ComponentModel.ISupportInitialize)(this.pictureBox_qrcode)).BeginInit();
            this.SuspendLayout();
            // 
            // pictureBox_qrcode
            // 
            this.pictureBox_qrcode.BorderStyle = System.Windows.Forms.BorderStyle.FixedSingle;
            this.pictureBox_qrcode.Location = new System.Drawing.Point(12, 12);
            this.pictureBox_qrcode.Name = "pictureBox_qrcode";
            this.pictureBox_qrcode.Size = new System.Drawing.Size(256, 256);
            this.pictureBox_qrcode.TabIndex = 0;
            this.pictureBox_qrcode.TabStop = false;
            // 
            // label_content
            // 
            this.label_content.AutoSize = true;
            this.label_content.Font = new System.Drawing.Font("Calibri", 9F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(0)));
            this.label_content.Location = new System.Drawing.Point(9, 271);
            this.label_content.Name = "label_content";
            this.label_content.Size = new System.Drawing.Size(102, 14);
            this.label_content.TabIndex = 1;
            this.label_content.Text = "Please input URL:";
            // 
            // textBox_content
            // 
            this.textBox_content.Font = new System.Drawing.Font("Calibri", 9F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(0)));
            this.textBox_content.Location = new System.Drawing.Point(12, 288);
            this.textBox_content.Name = "textBox_content";
            this.textBox_content.Size = new System.Drawing.Size(256, 22);
            this.textBox_content.TabIndex = 2;
            // 
            // button_generate
            // 
            this.button_generate.Font = new System.Drawing.Font("Calibri", 9F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(0)));
            this.button_generate.Location = new System.Drawing.Point(12, 358);
            this.button_generate.Name = "button_generate";
            this.button_generate.Size = new System.Drawing.Size(75, 23);
            this.button_generate.TabIndex = 3;
            this.button_generate.Text = "Generate";
            this.button_generate.UseVisualStyleBackColor = true;
            this.button_generate.Click += new System.EventHandler(this.Button_generate_Click);
            // 
            // button_clear
            // 
            this.button_clear.Font = new System.Drawing.Font("Calibri", 9F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(0)));
            this.button_clear.Location = new System.Drawing.Point(93, 358);
            this.button_clear.Name = "button_clear";
            this.button_clear.Size = new System.Drawing.Size(75, 23);
            this.button_clear.TabIndex = 4;
            this.button_clear.Text = "Clear";
            this.button_clear.UseVisualStyleBackColor = true;
            this.button_clear.Click += new System.EventHandler(this.Button_clear_Click);
            // 
            // button_select
            // 
            this.button_select.Font = new System.Drawing.Font("Calibri", 9F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(0)));
            this.button_select.Location = new System.Drawing.Point(193, 358);
            this.button_select.Name = "button_select";
            this.button_select.Size = new System.Drawing.Size(75, 23);
            this.button_select.TabIndex = 5;
            this.button_select.Text = "Select";
            this.button_select.UseVisualStyleBackColor = true;
            this.button_select.Click += new System.EventHandler(this.Button_select_Click);
            // 
            // textBox_images
            // 
            this.textBox_images.Font = new System.Drawing.Font("Calibri", 9F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(0)));
            this.textBox_images.Location = new System.Drawing.Point(12, 330);
            this.textBox_images.Name = "textBox_images";
            this.textBox_images.Size = new System.Drawing.Size(256, 22);
            this.textBox_images.TabIndex = 7;
            // 
            // label_path
            // 
            this.label_path.AutoSize = true;
            this.label_path.Font = new System.Drawing.Font("Calibri", 9F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(0)));
            this.label_path.Location = new System.Drawing.Point(9, 313);
            this.label_path.Name = "label_path";
            this.label_path.Size = new System.Drawing.Size(135, 14);
            this.label_path.TabIndex = 6;
            this.label_path.Text = "Please input save path:";
            // 
            // FormMain
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 12F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(280, 393);
            this.Controls.Add(this.textBox_images);
            this.Controls.Add(this.label_path);
            this.Controls.Add(this.button_select);
            this.Controls.Add(this.button_clear);
            this.Controls.Add(this.button_generate);
            this.Controls.Add(this.textBox_content);
            this.Controls.Add(this.label_content);
            this.Controls.Add(this.pictureBox_qrcode);
            this.Icon = ((System.Drawing.Icon)(resources.GetObject("$this.Icon")));
            this.Name = "FormMain";
            this.Text = "QRCode";
            this.FormClosed += new System.Windows.Forms.FormClosedEventHandler(this.FormMain_FormClosed);
            this.Load += new System.EventHandler(this.FormMain_Load);
            ((System.ComponentModel.ISupportInitialize)(this.pictureBox_qrcode)).EndInit();
            this.ResumeLayout(false);
            this.PerformLayout();

        }

        #endregion

        private System.Windows.Forms.PictureBox pictureBox_qrcode;
        private System.Windows.Forms.Label label_content;
        private System.Windows.Forms.TextBox textBox_content;
        private System.Windows.Forms.Button button_generate;
        private System.Windows.Forms.Button button_clear;
        private System.Windows.Forms.Button button_select;
        private System.Windows.Forms.TextBox textBox_images;
        private System.Windows.Forms.Label label_path;
    }
}

